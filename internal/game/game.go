package game

import (
	"EndlessStairsGolang/internal/character"
	"EndlessStairsGolang/internal/stair"
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"EndlessStairsGolang/internal/leaderboard"

	"github.com/eiannone/keyboard"
)

// InputProvider abstracts input for testability.
type InputProvider interface {
	GetInput(ctx context.Context) (string, error)
}

// OutputProvider abstracts output for testability.
type OutputProvider interface {
	Println(a ...interface{})
	Printf(format string, a ...interface{})
	Print(a ...interface{})
}

// StdIO implements InputProvider and OutputProvider for real terminal.
type StdIO struct{}

func (s *StdIO) Println(a ...interface{})               { fmt.Println(a...) }
func (s *StdIO) Printf(format string, a ...interface{}) { fmt.Printf(format, a...) }
func (s *StdIO) Print(a ...interface{})                 { fmt.Print(a...) }
func (s *StdIO) GetInput(ctx context.Context) (string, error) {
	inputCh := make(chan string)
	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				inputCh <- ""
				return
			}
			if key == keyboard.KeyArrowLeft || char == 'l' || char == 'L' {
				inputCh <- "left"
				return
			}
			if key == keyboard.KeyArrowRight || char == 'r' || char == 'R' {
				inputCh <- "right"
				return
			}
		}
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case input := <-inputCh:
		return input, nil
	}
}

// Game represents the Endless Stairs game state.
type Game struct {
	Player      *character.Character         // The player character
	Input       InputProvider                // Handles input abstraction
	Output      OutputProvider               // Handles output abstraction
	Leaderboard leaderboard.LeaderboardStore // Leaderboard storage
}

// NewGame creates a new Game instance with default IO and leaderboard file storage.
func NewGame() *Game {
	return &Game{
		Input:       &StdIO{},
		Output:      &StdIO{},
		Leaderboard: &leaderboard.FileLeaderboardStore{Path: "leaderboard.txt"},
	}
}

// StartGame runs the main game loop using goroutines, channels, and context.
// It handles player setup, game loop, rendering, input, and scoring.
func (g *Game) StartGame() {
	reader := bufio.NewReader(os.Stdin)
	g.Output.Printf("Enter your character's name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		g.Output.Println("Error reading name:", err)
		return
	}
	name = strings.TrimSpace(name)
	g.Player = character.NewCharacter(name)
	g.Output.Printf("Welcome, %s! Let's climb the endless stairs.\n", g.Player.Name)

	rand.Seed(time.Now().UnixNano())
	positions := []string{"left", "right"}
	g.Player.Position = positions[rand.Intn(2)] // Start on left or right stair

	if err := keyboard.Open(); err != nil {
		g.Output.Println("Failed to initialize keyboard input:", err)
		return
	}
	defer keyboard.Close()

	frameHeight := 14
	blankLine := strings.Repeat(" ", 22)

	for {
		timeLimit := getTimeLimit(g.Player.Score)
		clearScreen()
		g.Output.Println("Endless Stairs! (Use ←/→ arrows or l/r keys)")

		// Generate the next stair and its direction
		nextStair, nextDir, nextStairLines := generateNextStair()
		// If falling stair, halve the time limit but not below 0.9 seconds
		if nextStair.Type == stair.StairFalling {
			half := timeLimit / 2
			minLimit := time.Duration(0.9 * float64(time.Second))
			if half < minLimit {
				timeLimit = minLimit
			} else {
				timeLimit = half
			}
		}
		// If reverse polarity stair, swap left/right for this round
		reversePolarity := nextStair.Type == stair.StairReverse
		if reversePolarity {
			g.Output.Println("!! REVERSE POLARITY !! Left is right, right is left!")
		}
		// Build and print the frame
		frame := g.renderFrame(frameHeight, blankLine, nextStairLines)
		g.printFrame(frame)

		g.Output.Printf("\nScore: %d\n", g.Player.Score)
		g.Output.Printf("You have %.1f seconds to choose!\n", timeLimit.Seconds())
		g.Output.Printf("Jump left or right? (l/r): ")

		// Get player move with timeout
		move, err := g.getPlayerMove(timeLimit)
		if reversePolarity {
			if move == "left" {
				move = "right"
			} else if move == "right" {
				move = "left"
			}
		}
		if err != nil {
			g.Output.Println("\nTime's up! You fell!")
			g.Output.Printf("Game over, %s! Final score: %d\n", g.Player.Name, g.Player.Score)
			return
		}

		if move == nextDir {
			switch nextStair.Type {
			case stair.StairFalling:
				g.Output.Println("Oh no! The stair collapsed! You barely made it!")
				time.Sleep(400 * time.Millisecond)
				// Do not increase score, do not end game
			case stair.StairSpiked:
				if g.Player.Score > 0 {
					g.Player.Score--
				}
				g.Output.Println("Ouch! You landed on spiked stairs! Score -1.")
				time.Sleep(400 * time.Millisecond)
				// Continue the game
			case stair.StairSuper:
				g.Player.Score += 5
				g.Output.Println("Super stair! +5 points!")
				time.Sleep(400 * time.Millisecond)
			default:
				g.Player.Score++
				g.Output.Println("Good jump!")
				time.Sleep(200 * time.Millisecond)
			}
		} else {
			g.Output.Printf("Oops! The stair was to the %s. Game over, %s! Final score: %d\n", nextDir, g.Player.Name, g.Player.Score)
			break
		}
	}
}

// generateNextStair randomly picks the next stair direction and type, and returns the stair, its direction, and rendered lines.
func generateNextStair() (*stair.Stair, string, []string) {
	directions := []string{"left", "right"}
	dir := directions[rand.Intn(2)]
	// 50% normal, 10% falling, 10% spiked, 20% reverse, 10% super
	typeRoll := rand.Intn(100)
	var stairType string
	switch {
	case typeRoll < 50:
		stairType = stair.StairNormal
	case typeRoll < 60:
		stairType = stair.StairFalling
	case typeRoll < 70:
		stairType = stair.StairSpiked
	case typeRoll < 90:
		stairType = stair.StairReverse
	default:
		stairType = stair.StairSuper
	}
	st := stair.NewStair(dir, stairType)
	if dir == "left" {
		return st, dir, st.LeftRender()
	}
	return st, dir, st.RightRender()
}

// renderFrame builds the frame for the current step, padding with blank lines and adding the stair.
func (g *Game) renderFrame(frameHeight int, blankLine string, nextStairLines []string) []string {
	frame := make([]string, 0, frameHeight)
	for len(frame) < frameHeight-9 {
		frame = append(frame, blankLine)
	}
	frame = append(frame, nextStairLines...)
	return frame
}

// printFrame outputs the frame to the OutputProvider, line by line.
func (g *Game) printFrame(frame []string) {
	for _, line := range frame {
		g.Output.Println(line)
	}
}

// getPlayerMove handles input with a timeout and returns the move direction ("left" or "right").
// Returns an error if the time limit is exceeded.
func (g *Game) getPlayerMove(timeLimit time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()
	input, err := g.Input.GetInput(ctx)
	if err != nil {
		return "", err
	}
	switch input {
	case "left", "l":
		return "left", nil
	case "right", "r":
		return "right", nil
	default:
		return "", nil
	}
}

// getNextDirectionStair returns the next stair direction, generating one if needed.
func getNextDirectionStair(stairs []*stair.Stair, i int, directions []string) string {
	if i+1 < len(stairs) {
		return stairs[i+1].Direction
	}
	if rand.Intn(2) == 0 {
		return "left"
	}
	return "right"
}

// getTimeLimit returns the time limit based on the score.
func getTimeLimit(score int) time.Duration {
	// Base time in seconds (e.g., 10s at score 0)
	baseTime := 5.0

	// Time reduction per score point
	reductionPerPoint := 0.5

	// Calculate reduced time
	reducedTime := baseTime - (float64(score) * reductionPerPoint)

	// Ensure it's at least 1 second
	if reducedTime < 1.0 {
		reducedTime = 1.0
	}

	return time.Duration(reducedTime * float64(time.Second))
}

// clearScreen clears the terminal screen.
func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

// ShowStartMenu displays the start menu and handles user selection.
func (g *Game) ShowStartMenu() {
	for {
		clearScreen()
		g.Output.Println("==== Endless Stairs ====")
		g.Output.Println("1. Start Game")
		g.Output.Println("2. View Leaderboard")
		g.Output.Println("3. Quit")
		g.Output.Print("Select an option (1-3): ")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			g.StartGame()
			g.saveScoreToLeaderboard()
			g.Output.Println("Press Enter to return to menu...")
			reader.ReadString('\n')
		case "2":
			g.ShowLeaderboard()
			g.Output.Println("Press Enter to return to menu...")
			reader.ReadString('\n')
		case "3":
			g.Output.Println("Goodbye!")
			return
		default:
			g.Output.Println("Invalid option. Please try again.")
		}
	}
}

// ShowLeaderboard displays the leaderboard using the LeaderboardStore.
func (g *Game) ShowLeaderboard() {
	clearScreen()
	g.Output.Println("==== Leaderboard ====")
	entries, err := g.Leaderboard.TopN(10)
	if err != nil || len(entries) == 0 {
		g.Output.Println("No scores yet.")
		return
	}
	for i, entry := range entries {
		g.Output.Printf("%d. %s - %d\n", i+1, entry.Name, entry.Score)
	}
}

// saveScoreToLeaderboard saves the current player's score to the leaderboard.
func (g *Game) saveScoreToLeaderboard() {
	if g.Player == nil || g.Player.Score == 0 {
		return
	}
	entries, err := g.Leaderboard.Load()
	if err != nil {
		entries = nil
	}
	entries = append(entries, leaderboard.LeaderboardEntry{Name: g.Player.Name, Score: g.Player.Score})
	// Sort descending by score
	for i := 0; i < len(entries); i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[j].Score > entries[i].Score {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}
	if len(entries) > 10 {
		entries = entries[:10]
	}
	_ = g.Leaderboard.Save(entries)
}
