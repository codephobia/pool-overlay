package apidocs

import "github.com/codephobia/pool-overlay/libs/go/models"

// swagger:route GET /players players Players
// An array of existing players in the database.
// responses:
//   200: playersResp
//   422: errorResp
//   500: errorResp

// swagger:parameters Players
type PlayersParam struct {
	// The page of players to return.
	//
	// in: query
	// required: false
	// min: 1
	// default: 1
	// example: 1
	Page int `json:"page"`
}

// A page of players in the database.
// swagger:response playersResp
type PlayersRespWrapper struct {
	// in: body
	// maxItems: 10
	Body []models.Player
}
