package apidocs

import "github.com/codephobia/pool-overlay/libs/go/models"

// swagger:route GET /game game Game
// Get the current state of the game.
// responses:
//   200: gameResp

// The current state of the game.
// swagger:response gameResp
type GameRespWrapper struct {
	// in: body
	Body models.Game
}
