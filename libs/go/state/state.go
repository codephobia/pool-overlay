package state

import (
	"sync"

	"github.com/codephobia/pool-overlay/libs/go/models"
	"gorm.io/gorm"
)

type State struct {
	Game   *models.Game `json:"game"`
	Hidden bool         `json:"hidden"`

	mutex sync.Mutex
}

func NewState(db *gorm.DB) *State {
	return &State{
		Game:   models.NewGame(db).LoadLatest(),
		Hidden: true,
	}
}

func (s *State) ToggleHidden() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Hidden = !s.Hidden

	return s.Hidden
}
