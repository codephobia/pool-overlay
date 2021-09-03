package api

import (
	"log"
	"net/http"

	"github.com/codephobia/pool-overlay/libs/go/events"
	"github.com/codephobia/pool-overlay/libs/go/overlay"
)

// OverlayToggleResp is the response from toggling the overlay.
type OverlayToggleResp struct {
	Hidden bool `json:"hidden"`
}

// Handler for overlay.
func (server *Server) handleOverlay() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay handler for GET method.
func (server *Server) handleOverlayGet(w http.ResponseWriter, r *http.Request) {
	// upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR] unable to upgrade connection: %s", err)
		return
	}

	// create new connection for overlay
	oc := overlay.NewOverlayConnection(server.overlay, conn)
	if err != nil {
		log.Printf("[ERROR] overlay connect: %s", err)
		return
	}

	// register connection on overlay
	server.overlay.Register <- oc

	// init read / write for socket connection
	go oc.WritePump()
	go oc.ReadPump()
}

// Handler for overlay toggle.
func (server *Server) handleOverlayToggle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayToggleGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay toggle handler for GET method.
func (server *Server) handleOverlayToggleGet(w http.ResponseWriter, r *http.Request) {
	hidden := server.state.ToggleHidden()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayToggleEventType,
		events.NewOverlayToggleEventPayload(hidden),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	server.handleSuccess(w, r, &OverlayToggleResp{
		Hidden: hidden,
	})
}
