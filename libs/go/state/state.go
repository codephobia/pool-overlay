package state

import (
	"github.com/codephobia/pool-overlay/libs/go/models"
	"gorm.io/gorm"
)

// State is the current state of the overlay.
type State struct {
	Table   int           `json:"table"`
	Game    *models.Game  `json:"game"`
	Overlay *OverlayState `json:"overlay"`
}

// NewState creates a new state.
func NewState(db *gorm.DB, table int) *State {
	return &State{
		Table:   table,
		Game:    models.NewGame(db, table).LoadLatest(table),
		Overlay: NewOverlayState(table),
	}
}
