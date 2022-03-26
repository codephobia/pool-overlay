package apidocs

import (
	"github.com/codephobia/pool-overlay/libs/go/api"
)

// swagger:route GET /games/count games GamesCount
// The number of completed games in the database.
// responses:
//   200: gamesCountResp
//   500: errorResp

// The number of completed games in the database.
// swagger:response gamesCountResp
type GamesCountRespWrapper struct {
	// in: body
	Body api.CountResp
}
