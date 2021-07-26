package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route PATCH /game/race-to game GameRaceTo
// Update the race to amount of the game on the game state.
// responses:
//   200: gameRaceToResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameRaceTo
type GameUpdateRaceToParam struct {
	// The direction to change the amount by.
	//
	// in: body
	// required: true
	// example: increment
	Direction string `json:"direction"`
}

// The updated game race to amount.
// swagger:response gameRaceToResp
type GameRaceToRespWrapper struct {
	// in: body
	Body api.GameRaceToResp
}
