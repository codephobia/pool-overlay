package apidocs

// swagger:route PATCH /game/players game GamePlayers
// Sets a player number to a player ID.
// responses:
//   200: gameResp
//   404: errorResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GamePlayers
type GamePlayersParam struct {
	// The player number to set a player on.
	//
	// in: body
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`

	// The ID of the player to set on the playerNum.
	//
	// in: body
	// required: true
	// min: 1
	// example: 1
	PlayerID int `json:"playerID"`
}

// swagger:route DELETE /game/players game GamePlayersDelete
// Unsets a player number.
// responses:
//   200: gameResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GamePlayersDelete
type GamePlayersDeleteParam struct {
	// The player number to set a player on.
	//
	// in: body
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`
}
