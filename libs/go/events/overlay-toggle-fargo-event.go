package events

const OverlayToggleFargoEventType = "OVERLAY_TOGGLE_FARGO"

type OverlayToggleFargoEventPayload struct {
	ShowFargo bool `json:"showFargo"`
}

func NewOverlayToggleFargoEventPayload(showFargo bool) *OverlayToggleFargoEventPayload {
	return &OverlayToggleFargoEventPayload{
		ShowFargo: showFargo,
	}
}
