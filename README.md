# EndlessStairsGolang


Endless Stairs is a terminal-based game written in Go. Jump left or right to climb the endless stairs. The game gets faster as your score increases!

## How to Play
- Use ←/→ arrow keys or l/r keys to jump to the next stair.
- The next stair is shown beside your character.
- The game gets faster as your score increases.
- The game ends if you jump in the wrong direction or run out of time.

## Run the Game
```sh
go run ./cmd/main.go
```

## Requirements
- Go 1.18 or newer
- Terminal that supports ANSI escape codes (for screen clearing)

## Stair Types

- **Normal Stair**: Standard stair. Jump in the correct direction to score a point.
- **Falling Stair**: The stair collapses if you land on it! You barely make it, but your score does not increase. The time to select is halved (but never below 0.9 seconds).
- **Spiked Stair**: Ouch! Landing on this stair reduces your score by 1 (but not below 0). The game continues.
- **Reverse Polarity Stair**: Controls are reversed for this stair! Left becomes right and right becomes left for this round. Watch for the warning message!
- **Super Stair**: Jumping on this stair gives you a big bonus! Score +5 points. Watch for the special ASCII art ([***]).

Each stair type is visually distinct in the ASCII art.


