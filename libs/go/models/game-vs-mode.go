package models

// GameVsMode is the type of players to be used.
type GameVsMode int

const (
	// OneVsOne is one player vs one player.
	OneVsOne GameVsMode = iota
	// ScotchDoubles is a team of two players vs a team of two players.
	ScotchDoubles
	// Unexported end value used for validation.
	vsModeEnd
)

// Human readable versions of the GameVsModes.
var modeNames = map[GameVsMode]string{
	OneVsOne:      "1v1",
	ScotchDoubles: "Scotch Doubles",
}

// String returns the human readable version.
func (m GameVsMode) String() string {
	return modeNames[m]
}

// IsValid returns if the GameVsMode is valid.
func (m GameVsMode) IsValid() bool {
	return m < vsModeEnd
}
