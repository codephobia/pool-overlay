package api

import (
	"net/http"

	"github.com/codephobia/pool-overlay/libs/go/models"
)

// Handler for /flags.
func (server *Server) handleFlags() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleFlagsGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Flags handler for GET method.
func (server *Server) handleFlagsGet(w http.ResponseWriter, r *http.Request) {
	var flags []*models.Flag

	flagsResult := server.db.
		Select("id", "country", "image_path").
		Order("id").
		Find(&flags)
	if flagsResult.Error != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	// return results
	server.handleSuccess(w, r, flags)
}
