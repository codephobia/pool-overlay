package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route GET /game/update/race-to game GameUpdateRaceTo
// Update the race to amount of the game on the game state.
// responses:
//   200: gameRaceToResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameUpdateRaceTo
type GameUpdateRaceToParam struct {
	// The direction to change the amount by.
	//
	// in: query
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
