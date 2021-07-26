package apidocs

// swagger:route PATCH /game/players/flag game GamePlayersFlag
// Sets the flag for the specified player number.
// responses:
//   200: gameResp
//   404: errorResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GamePlayersFlag
type GamePlayersFlagParam struct {
	// The player number to set a player on.
	//
	// in: body
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`

	// The ID of the flag to set on the playerNum.
	//
	// in: body
	// required: true
	// min: 1
	// example: 1
	FlagID int `json:"flagID"`
}
