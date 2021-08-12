package apidocs

import "github.com/codephobia/pool-overlay/libs/go/models"

// swagger:route GET /flags flags Flags
// An array of existing flags in the database.
// responses:
//   200: flagsResp
//   500: errorResp

// Flags in the database.
// swagger:response flagsResp
type FlagsRespWrapper struct {
	// in: body
	Body []models.Flag
}
