package state

import "sync"

type OverlayState struct {
	Table                     int  `json:"table"`
	Hidden                    bool `json:"hidden"`
	ShowFlags                 bool `json:"showFlags"`
	ShowFargo                 bool `json:"showFargo"`
	ShowScore                 bool `json:"showScore"`
	WaitingForPlayers         bool `json:"waitingForPlayers"`
	WaitingForTournamentStart bool `json:"waitingForTournamentStart"`
	TableNoLongerInUse        bool `json:"tableNoLongerInUse"`

	mutex sync.Mutex `json:"-"`
}

// NewOverlayState creates a new overlay state.
func NewOverlayState(table int) *OverlayState {
	return &OverlayState{
		Table:     table,
		Hidden:    false,
		ShowFlags: true,
		ShowFargo: true,
		ShowScore: true,
	}
}

// SetHidden sets the visibility of the overlay.
func (s *OverlayState) SetHidden(hidden bool) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Hidden = hidden

	return s.Hidden
}

// ToggleHidden toggles the visibility of the overlay.
func (s *OverlayState) ToggleHidden() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Hidden = !s.Hidden

	return s.Hidden
}

// SetFlags sets the visibility of flags on the overlay.
func (s *OverlayState) SetFlags(showFlags bool) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowFlags = showFlags

	return s.ShowFlags
}

// ToggleFlags toggles the visibility of flags on the overlay.
func (s *OverlayState) ToggleFlags() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowFlags = !s.ShowFlags

	return s.ShowFlags
}

// SetFargo sets the visibility of fargo ratings on the overlay.
func (s *OverlayState) SetFargo(showFargo bool) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowFargo = showFargo

	return s.ShowFargo
}

// ToggleFargo toggles the visibility of fargo ratings on the overlay.
func (s *OverlayState) ToggleFargo() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowFargo = !s.ShowFargo

	return s.ShowFargo
}

// SetScore sets the visibility of the player scores on the overlay.
func (s *OverlayState) SetScore(showScore bool) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowScore = showScore

	return s.ShowScore
}

// ToggleScore toggles the visibility of the player scores on the overlay.
func (s *OverlayState) ToggleScore() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowScore = !s.ShowScore

	return s.ShowScore
}
