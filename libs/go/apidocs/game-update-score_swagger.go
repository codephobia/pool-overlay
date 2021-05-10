package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route GET /game/update/score game GameUpdateScore
// Increment / decrement the score of a player.
// responses:
//   200: gameScoreResp
//   422: errorResp
//   500: errorResp

// swagger:parameters GameUpdateScore
type GameUpdateScoreParam struct {
	// The player number of the score being changed.
	//
	// in: query
	// required: true
	// min: 1
	// max: 2
	// example: 1
	PlayerNum int `json:"playerNum"`

	// The direction to change the score.
	//
	// in: query
	// required: true
	// example: increment
	Direction string `json:"direction"`
}

// The newly updated score of the game.
// swagger:response gameScoreResp
type GameScoreRespWrapper struct {
	// in: body
	Body api.GameScoreResp
}
