package models

// GameType is the type of pool gaming being played.
type GameType uint

const (
	// EightBall is an 8 ball game.
	EightBall GameType = iota
	// NineBall is a 9 ball game.
	NineBall
	// TenBall is a 10 ball game.
	TenBall
	// Unexported end value used for validation.
	typeEnd
)

// Human readable versions of the Types.
var typeNames = map[GameType]string{
	EightBall: "8 Ball",
	NineBall:  "9 Ball",
	TenBall:   "10 Ball",
}

// String returns the human readable version.
func (m GameType) String() string {
	return typeNames[m]
}

// IsValid returns if the GameType is valid.
func (m GameType) IsValid() bool {
	return m < typeEnd
}
