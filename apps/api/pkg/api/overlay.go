package api

import (
	"log"
	"net/http"

	"github.com/codephobia/pool-overlay/apps/api/pkg/overlay"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
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
