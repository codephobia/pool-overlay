package apidocs

import "github.com/codephobia/pool-overlay/libs/go/models"

// swagger:route GET /players players PlayersGet
// An array of existing players in the database.
// responses:
//   200: playersGetResp
//   422: errorResp
//   500: errorResp

// swagger:parameters PlayersGet
type PlayersGetParam struct {
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
// swagger:response playersGetResp
type PlayersGetRespWrapper struct {
	// in: body
	// maxItems: 10
	Body []models.Player
}

// swagger:route POST /players players PlayersPost
// Add a new player to the database.
// responses:
//   200: playersPostResp
//   422: errorResp
//   500: errorResp

// PlayerPostData is the body content for a new player.
type PlayerPostData struct {
	// The new player name.
	//
	// in: body
	// required: true
	// example: Joe
	Name string `json:"name"`

	// The new player flag id.
	//
	// in: body
	// required: true
	// example: 1
	FlagID uint `json:"flag_id"`

	// The new player Fargo id.
	//
	// in: body
	// required: true
	// example: 1
	FargoID uint `json:"fargo_id"`

	// The new player Fargo rating.
	//
	// in: body
	// required: true
	// example: 1
	FargoRating uint `json:"fargo_rating"`
}

// swagger:parameters PlayersPost
type PlayersPostParam struct {
	// The player details.
	//
	// in: body
	// required: true
	// example: { "name": "Joe", "flag_id": 1, "fargo_id": 1, "fargo_rating": 1 }
	Player PlayerPostData
}

// The newly created player with id.
// swagger:response playersPostResp
type PlayersPostRespWrapper struct {
	// in: body
	Body models.Player
}
