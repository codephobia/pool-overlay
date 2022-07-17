package state

import (
	"sync"

	"github.com/codephobia/pool-overlay/libs/go/models"
	"gorm.io/gorm"
)

// State is the current state of the overlay.
type State struct {
	Game      *models.Game `json:"game"`
	Hidden    bool         `json:"hidden"`
	ShowFlags bool         `json:"showFlags"`
	ShowFargo bool         `json:"showFargo"`
	ShowScore bool         `json:"showScore"`

	mutex sync.Mutex
}

// NewState creates a new state.
func NewState(db *gorm.DB) *State {
	return &State{
		Game:      models.NewGame(db).LoadLatest(),
		Hidden:    false,
		ShowFlags: true,
		ShowFargo: true,
		ShowScore: true,
	}
}

// ToggleHidden toggles the visibility of the overlay.
func (s *State) ToggleHidden() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Hidden = !s.Hidden

	return s.Hidden
}

// ToggleFlags toggles the visibility of flags on the overlay.
func (s *State) ToggleFlags() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowFlags = !s.ShowFlags

	return s.ShowFlags
}

// ToggleFargo toggles the visibility of fargo ratings on the overlay.
func (s *State) ToggleFargo() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowFargo = !s.ShowFargo

	return s.ShowFargo
}

// ToggleScore toggles the visibility of the player scores on the overlay.
func (s *State) ToggleScore() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowScore = !s.ShowScore

	return s.ShowScore
}
