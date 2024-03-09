package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// swagger:route GET /table/{tableNum}/overlay/toggle overlay OverlayToggle
// Toggles showing / hiding of the overlay.
// responses:
//   200: overlayToggleResp
//   422: errorResp

// swagger:parameters OverlayToggle
type OverlayToggleGetParam struct {
	// The table number to toggle the overlay on.
	//
	// in: path
	// required: true
	// example: 1
	TableNum int `json:"tableNum"`
}

// nolint
// Returns updated value for Hidden.
// swagger:response overlayToggleResp
type overlayToggleRespWrapper struct {
	// in: body
	Body api.OverlayToggleResp
}
