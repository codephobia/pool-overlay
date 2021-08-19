package apidocs

import "github.com/codephobia/pool-overlay/libs/go/models"

// swagger:route GET /players/{playerID} players PlayersByIdGet
// Get a player by their ID.
// responses:
//   200: playersByIdGetResp
//   404: errorResp
//   422: errorResp
//   500: errorResp

// swagger:parameters PlayersByIdGet
type PlayersByIdGetParam struct {
	// The ID of the player to return.
	//
	// in: path
	// required: true
	// example: 1
	PlayerID int `json:"playerID"`
}

// The player with the specified ID.
// swagger:response playersByIdGetResp
type PlayersByIdGetRespWrapper struct {
	// in: body
	Body models.Player
}

// swagger:route PATCH /players/{playerID} players PlayersByIdPatch
// Update a player by their ID.
// responses:
//   200: playersByIdPatchResp
//   404: errorResp
//   422: errorResp
//   500: errorResp

// PlayerPatchData is the body content for a player update.
type PlayerPatchData struct {
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
}

// swagger:parameters PlayersByIdPatch
type PlayersByIdPatchParam struct {
	// The ID of the player to update.
	//
	// in: path
	// required: true
	// example: 1
	PlayerID int `json:"playerID"`

	// The player data to be updated.
	//
	// in: body
	// required: true
	// example: { "name": "Joe", "flag_id": 1 }
	Player PlayerPatchData
}

// The updated player details.
// swagger:response playersByIdPatchResp
type PlayersByIdPatchRespWrapper struct {
	// in: body
	Body models.Player
}

// swagger:route DELETE /players/{playerID} players PlayersByIdDelete
// Deletes a player by their ID.
// responses:
//   204: playersByIdDeleteResp
//   404: errorResp
//   422: errorResp
//   500: errorResp

// swagger:parameters PlayersByIdDelete
type PlayersByIdDeleteParam struct {
	// The ID of the player to delete.
	//
	// in: path
	// required: true
	// example: 1
	PlayerID int `json:"playerID"`
}

// An empty 204 No Content response.
// swagger:response playersByIdDeleteResp
type PlayersByIdDeleteRespWrapper struct{}
