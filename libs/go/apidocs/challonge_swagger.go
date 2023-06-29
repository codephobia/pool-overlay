package apidocs

import (
	"github.com/codephobia/pool-overlay/libs/go/challonge"
)

// swagger:route GET /tournaments tournaments Tournaments
// An array of unfinished tournaments in Challonge.
// responses:
//   200: tournamentsResp
//   404: errorResp

// Incomplete tournaments in Challonge.
// swagger:response tournamentsResp
type TournamentsRespWrapper struct {
	// in: body
	Body []challonge.Tournament
}
