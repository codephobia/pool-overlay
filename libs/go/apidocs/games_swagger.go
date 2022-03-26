package apidocs

import "github.com/codephobia/pool-overlay/libs/go/models"

// swagger:route GET /games games GamesGet
// An array of existing completed games in the database.
// responses:
//   200: gamesGetResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GamesGet
type GamesGetParam struct {
	// The page of games to return.
	//
	// in: query
	// required: false
	// min: 1
	// default: 1
	// example: 1
	Page int `json:"page"`
}

// A page of completed games in the database.
// swagger:response gamesGetResp
type GamesGetRespWrapper struct {
	// in: body
	// maxItems: 10
	Body []models.Game
}
