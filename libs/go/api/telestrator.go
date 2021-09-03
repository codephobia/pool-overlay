package api

import (
	"log"
	"net/http"

	"github.com/codephobia/pool-overlay/libs/go/telestrator"
)

// Handler for telestrator.
func (server *Server) handleTelestrator() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleTelestratorGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Telestrator handler for GET method.
func (server *Server) handleTelestratorGet(w http.ResponseWriter, r *http.Request) {
	// upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR] unable to upgrade connection: %s", err)
		return
	}

	// create new connection for overlay
	oc := telestrator.NewOverlayConnection(server.telestrator.Overlay, conn)
	if err != nil {
		log.Printf("[ERROR] overlay connect: %s", err)
		return
	}

	// register connection on overlay
	server.telestrator.Overlay.Register <- oc

	// init read / write for socket connection
	go oc.WritePump()
	go oc.ReadPump()
}
