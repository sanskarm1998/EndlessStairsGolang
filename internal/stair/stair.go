package stair

// StairType constants
const (
	StairNormal  = "normal"
	StairFalling = "falling"
	StairSpiked  = "spiked"
	StairReverse = "reverse" // reverse polarity stairs
	StairSuper   = "super"   // super stair
	// Power-up and collectible stairs
	StairSlowTime = "slowtime"   // slows time for next N stairs
	StairDouble   = "double"     // double points for next N stairs
	StairShield   = "shield"     // protects from next hazard
	StairCoin     = "coin"       // collectible coin
	StairGem      = "gem"        // collectible gem
)

// Stair represents a stair in the Endless Stairs game.
type Stair struct {
	Direction string // "left" or "right"
	Type      string // "normal", "falling", "spiked"
}

// NewStair creates a new Stair with the given direction and type.
func NewStair(direction, stairType string) *Stair {
	return &Stair{Direction: direction, Type: stairType}
}

// LeftRender returns the stair ASCII art at the left offset.
func (s *Stair) LeftRender() []string {
	switch s.Type {
	case StairFalling:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[ v ] / \\ ",
			"     [___]",
		}
	case StairSpiked:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[^^^] / \\ ",
			"     [___]",
		}
	case StairReverse:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[<=>] / \\ ",
			"     [___]",
		}
	case StairSuper:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[***] / \\ ",
			"     [___]",
		}
	case StairSlowTime:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[SLOW] / \\ ",
			"     [___]",
		}
	case StairDouble:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[2X!!] / \\ ",
			"     [___]",
		}
	case StairShield:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[SHLD] / \\ ",
			"     [___]",
		}
	case StairCoin:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[ $ ] / \\ ",
			"     [___]",
		}
	case StairGem:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[<>] / \\ ",
			"     [___]",
		}
	default:
		return []string{
			"       O  ",
			"      /|\\ ",
			"[___] / \\ ",
			"     [___]",
		}
	}
}

// RightRender returns the stair ASCII art at the right offset.
func (s *Stair) RightRender() []string {
	switch s.Type {
	case StairFalling:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [ V ]",
			"    [___]  ",
		}
	case StairSpiked:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [^^^]",
			"    [___]      ",
		}
	case StairReverse:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [<=>]",
			"    [___]      ",
		}
	case StairSuper:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [***]",
			"    [___]      ",
		}
	case StairSlowTime:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [SLOW]",
			"    [___]      ",
		}
	case StairDouble:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [2X!!]",
			"    [___]      ",
		}
	case StairShield:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [SHLD]",
			"    [___]      ",
		}
	case StairCoin:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [ $ ]",
			"    [___]      ",
		}
	case StairGem:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [<>]",
			"    [___]      ",
		}
	default:
		return []string{
			"      O        ",
			"     /|\\       ",
			"     / \\ [___]",
			"    [___]      ",
		}
	}
} 