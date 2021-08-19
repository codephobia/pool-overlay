package apidocs

import (
	"github.com/codephobia/pool-overlay/libs/go/api"
)

// swagger:route GET /players/count players PlayersCount
// The number of players in the database.
// responses:
//   200: playersCountResp
//   500: errorResp

// The number of players in the database.
// swagger:response playersCountResp
type PlayersCountRespWrapper struct {
	// in: body
	Body api.CountResp
}
