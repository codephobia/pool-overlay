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

// handleOverlay
func (api *API) handleOverlay() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			api.handleOverlayGet(w, r)
		}
	})
}

// handleOverlayGet
func (api *API) handleOverlayGet(w http.ResponseWriter, r *http.Request) {
	// upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR] unable to upgrade connection: %s", err)
		return
	}

	// create new connection for overlay
	oc := overlay.NewOverlayConnection(api.overlay, conn)
	if err != nil {
		log.Printf("[ERROR] overlay connect: %s", err)
		return
	}

	// register connection on overlay
	api.overlay.Register <- oc

	// init read / write for socket connection
	go oc.WritePump()
	go oc.ReadPump()
}
