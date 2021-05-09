package apidocs

import "github.com/codephobia/pool-overlay/libs/go/api"

// Returns an error message.
// swagger:response errorResp
type errorRespWrapper struct {
	// in: body
	Body api.ErrorResp
}
