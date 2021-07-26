package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route PATCH /game/type game GameType
// Update the type of game on the game state.
// responses:
//   200: gameTypeResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameType
type GameTypeBody struct {
	// The type of game.
	//
	// in: body
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
