package apidocs

import (
	"github.com/codephobia/pool-overlay/libs/go/api"
)

// swagger:route GET /table/count table TableCount
// The number of current tables in use.
// responses:
//   200: tableCountResp

// The number of tables in use.
// swagger:response tableCountResp
type TableCountRespWrapper struct {
	// in: body
	Body api.CountResp
}
