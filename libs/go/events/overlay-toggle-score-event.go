package events

const OverlayToggleScoreEventType = "OVERLAY_TOGGLE_SCORE"

type OverlayToggleScoreEventPayload struct {
	ShowScore bool `json:"showScore"`
}

func NewOverlayToggleScoreEventPayload(showScore bool) *OverlayToggleScoreEventPayload {
	return &OverlayToggleScoreEventPayload{
		ShowScore: showScore,
	}
}
