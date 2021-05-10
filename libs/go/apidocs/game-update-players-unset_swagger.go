package apidocs

// swagger:route GET /game/update/players/unset game GameUpdatePlayersUnset
// Unsets a player number.
// responses:
//   200: gameResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameUpdatePlayersUnset
type GameUpdatePlayersUnsetParam struct {
	// The player number to set a player on.
	//
	// in: query
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`
}
