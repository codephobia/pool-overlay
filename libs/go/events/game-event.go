package events

import "github.com/codephobia/pool-overlay/libs/go/models"

const GameEventType = "GAME_EVENT"

type GameEventPayload struct {
	Game *models.Game `json:"game"`
}

func NewGameEventPayload(game *models.Game) *GameEventPayload {
	return &GameEventPayload{
		Game: game,
	}
}
