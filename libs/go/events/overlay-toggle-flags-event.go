package events

const OverlayToggleFlagsEventType = "OVERLAY_TOGGLE_FLAGS"

type OverlayToggleFlagsEventPayload struct {
	ShowFlags bool `json:"showFlags"`
}

func NewOverlayToggleFlagsEventPayload(showFlags bool) *OverlayToggleFlagsEventPayload {
	return &OverlayToggleFlagsEventPayload{
		ShowFlags: showFlags,
	}
}
