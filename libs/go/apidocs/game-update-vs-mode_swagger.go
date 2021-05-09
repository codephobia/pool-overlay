package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route GET /game/update/vs-mode game GameUpdateVsMode
// Update the vs-mode of game on the game state.
// responses:
//   200: gameVsModeResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameUpdateVsMode
type GameUpdateVsModeParam struct {
	// The vs-mode to change the game to.
	//
	// in: query
	// required: true
	// min: 0
	// max: 1
	// example: 0
	VsMode uint `json:"vs_mode"`
}

// The updated game vs-mode.
// swagger:response gameVsModeResp
type GameVsModeRespWrapper struct {
	// in: body
	Body api.GameVsModeResp
}
