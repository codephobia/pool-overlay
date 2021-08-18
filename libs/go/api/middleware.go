package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResp is an error response.
type ErrorResp struct {
	Err string `json:"error"`
}

// Index route handler. Since this is an API only server, we return a 401 if we
// try to access the root endpoint.
func (server *Server) handleIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.SetupResponse(w, r)

		// Handle preflight.
		if r.Method == "OPTIONS" {
			server.HandleOptions(w, r)
			return
		}

		server.handleError(w, r, http.StatusForbidden, ErrEndpointForbidden)
	})
}

// Healthcheck endpoint for Docker.
func (server *Server) handleHealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.SetupResponse(w, r)

		log.Printf("[HEALTHCHECK]")

		server.handleSuccess(w, r, "OK")
	})
}

// HandleSuccess handles a success response.
func (server *Server) handleSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	server.SetupResponse(w, r)

	// add headers to response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// return data
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		log.Printf("[WARN] unable to encode data: %s", err)
	}

	// close request
	(*r).Body.Close()
}

// handle204Success handles a success response with a 204 header.
func (server *Server) handle204Success(w http.ResponseWriter, r *http.Request) {
	server.SetupResponse(w, r)

	// add headers to response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	// close request
	(*r).Body.Close()
}

// HandleError handles an error response
func (server *Server) handleError(w http.ResponseWriter, r *http.Request, status int, err error) {
	server.SetupResponse(w, r)

	// add headers to response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	// return error message
	enc := json.NewEncoder(w)
	if err := enc.Encode(&ErrorResp{
		Err: err.Error(),
	}); err != nil {
		log.Printf("[WARN] unable to encode data: %s", err)
	}

	// close request
	(*r).Body.Close()
}
