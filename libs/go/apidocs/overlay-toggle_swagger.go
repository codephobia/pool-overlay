package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route GET /overlay/toggle overlay-toggle OverlayToggle
// Toggles showing / hiding of the overlay.
// responses:
//   200: overlayToggleResp
//   422: errorResp

// Returns updated value for Hidden.
// swagger:response overlayToggleResp
type overlayToggleRespWrapper struct {
	// in: body
	Body api.OverlayToggleResp
}
