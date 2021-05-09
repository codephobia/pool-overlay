package apidocs

import "github.com/codephobia/pool-overlay/libs/go/models"

// swagger:route GET /players/{playerID} players-by-id PlayersById
// Get a player by their ID.
// responses:
//   200: playersByIdResp
//   404: errorResp
//   422: errorResp
//   500: errorResp

// swagger:parameters PlayersById
type PlayersByIdParam struct {
	// The ID of the player to return.
	//
	// in: path
	// required: true
	// example: 1
	PlayerID int `json:"playerID"`
}

// The player with the specified ID.
// swagger:response playersByIdResp
type PlayersByIdRespWrapper struct {
	// in: body
	Body models.Player
}
