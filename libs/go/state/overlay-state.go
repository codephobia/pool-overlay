package state

import "sync"

type OverlayState struct {
	Hidden    bool `json:"hidden"`
	ShowFlags bool `json:"showFlags"`
	ShowFargo bool `json:"showFargo"`
	ShowScore bool `json:"showScore"`

	mutex sync.Mutex
}

// NewOverlayState creates a new overlay state.
func NewOverlayState() *OverlayState {
	return &OverlayState{
		Hidden:    false,
		ShowFlags: true,
		ShowFargo: true,
		ShowScore: true,
	}
}

// ToggleHidden toggles the visibility of the overlay.
func (s *OverlayState) ToggleHidden() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Hidden = !s.Hidden

	return s.Hidden
}

// ToggleFlags toggles the visibility of flags on the overlay.
func (s *OverlayState) ToggleFlags() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowFlags = !s.ShowFlags

	return s.ShowFlags
}

// ToggleFargo toggles the visibility of fargo ratings on the overlay.
func (s *OverlayState) ToggleFargo() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowFargo = !s.ShowFargo

	return s.ShowFargo
}

// ToggleScore toggles the visibility of the player scores on the overlay.
func (s *OverlayState) ToggleScore() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ShowScore = !s.ShowScore

	return s.ShowScore
}
