package api

import "net/http"

// SetupResponse adds CORS headers to an API response.
func (server *Server) SetupResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
}

// HandleOptions returns a 200 status code for preflight checks.
func (server *Server) HandleOptions(w http.ResponseWriter, r *http.Request) {
	server.SetupResponse(w, r)
	w.WriteHeader(http.StatusOK)
	(*r).Body.Close()
}
