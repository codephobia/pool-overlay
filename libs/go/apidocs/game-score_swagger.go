package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route PATCH /game/score game GameScorePatch
// Increment / decrement the score of a player.
// responses:
//   200: gameScoreResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameScorePatch
type GameScoreParam struct {
	// The player number of the score being changed.
	//
	// in: body
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`

	// The direction to change the score.
	//
	// in: body
	// required: true
	// example: increment
	Direction string `json:"direction"`
}

// swagger:route DELETE /game/score game GameScoreDelete
// Resets the score to zero for both players.
// responses:
//   200: gameScoreResp
//   422: errorResp
//   500: errorResp

// The newly updated score of the game.
// swagger:response gameScoreResp
type GameScoreRespWrapper struct {
	// in: body
	Body api.GameScoreResp
}
