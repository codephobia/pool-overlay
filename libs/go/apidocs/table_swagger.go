package apidocs

import (
	"github.com/codephobia/pool-overlay/libs/go/api"
)

// swagger:route POST /table/add table TableAdd
// Add a table to the current count.
// responses:
//   200: tableAddCountResp

// The number of tables in use.
// swagger:response tableAddCountResp
type TableAddRespWrapper struct {
	// in: body
	Body api.CountResp
}

// swagger:route POST /table/remove table TableRemove
// Removes a table from the current count.
// responses:
//   200: tableRemoveCountResp
//   422: errorResp

// The number of tables in use.
// swagger:response tableRemoveCountResp
type TableRemoveRespWrapper struct {
	// in: body
	Body api.CountResp
}
