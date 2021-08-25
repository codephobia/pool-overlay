package apidocs

import "github.com/codephobia/pool-overlay/libs/go/models"

// swagger:route GET /game game GameGet
// Get the current state of the game.
// responses:
//   200: gameGetResp

// The current state of the game.
// swagger:response gameGetResp
type GameGetRespWrapper struct {
	// in: body
	Body models.Game
}

// swagger:route POST /game game GamePost
// Mark the current game as completed and create a new one with the same players
// and settings.
// responses:
//   204: gamePostResp
//   422: errorResp
//   500: errorResp

// No content.
// swagger:response gamePostResp
type GamePostRespWrapper struct{}
