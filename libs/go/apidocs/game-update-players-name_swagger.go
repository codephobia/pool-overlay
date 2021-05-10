package apidocs

// swagger:route GET /game/update/players/name game GameUpdatePlayersName
// Sets the name for the specified player number.
// responses:
//   200: gameResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameUpdatePlayersName
type GameUpdatePlayersNameParam struct {
	// The player number to set a player on.
	//
	// in: query
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`

	// The action to apply to the character name.
	//
	// in: query
	// required: true
	// example: backspace
	Action string `json:"action"`
}
