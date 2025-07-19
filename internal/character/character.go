package character

// Character represents a player in the Endless Stairs game.
type Character struct {
	Name     string
	Score    int
	Position string // "left" or "right"
}

// NewCharacter creates a new Character with the given name, starting at left position.
func NewCharacter(name string) *Character {
	return &Character{
		Name:     name,
		Score:    0,
		Position: "left",
	}
}

