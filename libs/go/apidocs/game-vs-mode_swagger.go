package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route PATCH /game/vs-mode game GameVsMode
// Update the vs-mode of game on the game state.
// responses:
//   200: gameVsModeResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameVsMode
type GameUpdateVsModeParam struct {
	// The vs-mode to change the game to.
	//
	// in: body
	// required: true
	// min: 0
	// max: 1
	// example: 0
	VsMode uint `json:"vsMode"`
}

// The updated game vs-mode.
// swagger:response gameVsModeResp
type GameVsModeRespWrapper struct {
	// in: body
	Body api.GameVsModeResp
}
