package state

import "github.com/codephobia/pool-overlay/apps/api/pkg/models"

type State struct {
	Game   *models.Game `json:"game"`
	Hidden bool         `json:"hidden"`
}

func NewState() *State {
	return &State{
		Game:   models.NewGame(),
		Hidden: true,
	}
}
