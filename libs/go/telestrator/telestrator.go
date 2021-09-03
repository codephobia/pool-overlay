package telestrator

// Telestrator handles the overlay for telestration.
type Telestrator struct {
	Overlay *Overlay
}

// NewTelestrator returns a new Telestrator.
func NewTelestrator() *Telestrator {
	return &Telestrator{
		Overlay: NewOverlay(),
	}
}
