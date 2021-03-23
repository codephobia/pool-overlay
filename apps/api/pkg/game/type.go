package game

// Type is the type of pool gaming being played.
type Type uint

const (
	// EightBall is an 8 ball game.
	EightBall Type = iota
	// NineBall is a 9 ball game.
	NineBall
	// TenBall is a 10 ball game.
	TenBall
	// Unexported end value used for validation.
	typeEnd
)

// Human readable versions of the Types.
var typeNames = map[Type]string{
	EightBall: "8 Ball",
	NineBall:  "9 Ball",
	TenBall:   "10 Ball",
}

// String returns the human readable version.
func (m Type) String() string {
	return typeNames[m]
}

// IsValid returns if the Type is valid.
func (m Type) IsValid() bool {
	return m < typeEnd
}
