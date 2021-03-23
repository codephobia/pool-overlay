package game

// VsMode is the type of players to be used.
type VsMode int

const (
	// OneVsOne is one player vs one player.
	OneVsOne VsMode = iota
	// ScotchDoubles is a team of two players vs a team of two players.
	ScotchDoubles
	// Unexported end value used for validation.
	vsModeEnd
)

// Human readable versions of the VsModes.
var modeNames = map[VsMode]string{
	OneVsOne:      "1v1",
	ScotchDoubles: "Scotch Doubles",
}

// String returns the human readable version.
func (m VsMode) String() string {
	return modeNames[m]
}

// IsValid returns if the VsMode is valid.
func (m VsMode) IsValid() bool {
	return m < vsModeEnd
}
