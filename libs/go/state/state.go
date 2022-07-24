package state

import (
	"github.com/codephobia/pool-overlay/libs/go/models"
	"gorm.io/gorm"
)

// State is the current state of the overlay.
type State struct {
	Game    *models.Game  `json:"game"`
	Overlay *OverlayState `json:"overlay"`
}

// NewState creates a new state.
func NewState(db *gorm.DB) *State {
	return &State{
		Game:    models.NewGame(db).LoadLatest(),
		Overlay: NewOverlayState(),
	}
}
