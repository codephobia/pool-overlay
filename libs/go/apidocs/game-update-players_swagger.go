package apidocs

// swagger:route GET /game/update/players game GameUpdatePlayers
// Sets a player number to a player ID.
// responses:
//   200: gameResp
//   404: errorResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameUpdatePlayers
type GameUpdatePlayersParam struct {
	// The player number to set a player on.
	//
	// in: query
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`

	// The ID of the player to set on the playerNum.
	//
	// in: query
	// required: true
	// min: 1
	// example: 1
	PlayerID int `json:"playerID"`
}
