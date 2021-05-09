package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route GET /game/update/type game GameUpdateType
// Update the type of game on the game state.
// responses:
//   200: gameTypeResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameUpdateType
type GameUpdateTypeParam struct {
	// The type of game.
	//
	// in: query
	// required: true
	// min: 0
	// max: 2
	// example: 0
	Type uint `json:"type"`
}

// The updated game type.
// swagger:response gameTypeResp
type GameTypeRespWrapper struct {
	// in: body
	Body api.GameTypeResp
}
