package api

import (
	"encoding/json"
	"net/http"
)

// DataResp is an api response.
type DataResp struct {
	Data interface{} `json:"data"`
}

// HandleSuccess handles a success response
func (api *API) HandleSuccess(w *http.ResponseWriter, r *http.Request, data interface{}) {
	// add headers to response
	(*w).Header().Add("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusOK)

	// return data
	enc := json.NewEncoder(*w)
	enc.Encode(&DataResp{
		Data: data,
	})

	// close request
	(*r).Body.Close()
}
