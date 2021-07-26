package apidocs

// swagger:route PATCH /game/players/name game GamePlayersName
// Sets the name for the specified player number.
// responses:
//   200: gameResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GamePlayersName
type GamePlayersNameParam struct {
	// The player number to set a player on.
	//
	// in: body
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`

	// The new player name.
	//
	// in: body
	// required: true
	// example: John
	Name string `json:"name"`
}
