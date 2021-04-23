package events

const OverlayToggleEventType = "OVERLAY_TOGGLE"

type OverlayToggleEventPayload struct {
	Hidden bool `json:"hidden"`
}

func NewOverlayToggleEventPayload(hidden bool) *OverlayToggleEventPayload {
	return &OverlayToggleEventPayload{
		Hidden: hidden,
	}
}
