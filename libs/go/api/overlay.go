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

// OverlayToggleFlagsResp is the response from toggling flags on the overlay.
type OverlayToggleFlagsResp struct {
	ShowFlags bool `json:"showFlags"`
}

// OverlayToggleFargoResp is the response from toggling fargo ratings on the overlay.
type OverlayToggleFargoResp struct {
	ShowFargo bool `json:"showFargo"`
}

// OverlayToggleScoreResp is the response from toggling the player scores on the overlay.
type OverlayToggleScoreResp struct {
	ShowScore bool `json:"showScore"`
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

	// Generate current overlay state message.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[1].Overlay,
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Send current state to new connection.
	oc.Send <- message
}

// Handler for overlay toggle.
func (server *Server) handleOverlayToggle(table int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayToggleGet(w, r, table)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay toggle handler for GET method.
func (server *Server) handleOverlayToggleGet(w http.ResponseWriter, r *http.Request, table int) {
	hidden := server.tables[table].Overlay.ToggleHidden()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[table].Overlay,
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

// Handler for overlay toggle flags.
func (server *Server) handleOverlayToggleFlags(table int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayToggleFlagsGet(w, r, table)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay toggle flags handler for GET method.
func (server *Server) handleOverlayToggleFlagsGet(w http.ResponseWriter, r *http.Request, table int) {
	showFlags := server.tables[table].Overlay.ToggleFlags()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[table].Overlay,
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	server.handleSuccess(w, r, &OverlayToggleFlagsResp{
		ShowFlags: showFlags,
	})
}

// Handler for overlay toggle fargo.
func (server *Server) handleOverlayToggleFargo(table int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayToggleFargoGet(w, r, table)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay toggle fargo handler for GET method.
func (server *Server) handleOverlayToggleFargoGet(w http.ResponseWriter, r *http.Request, table int) {
	showFargo := server.tables[table].Overlay.ToggleFargo()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[table].Overlay,
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	server.handleSuccess(w, r, &OverlayToggleFargoResp{
		ShowFargo: showFargo,
	})
}

// Handler for overlay toggle score.
func (server *Server) handleOverlayToggleScore(table int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayToggleScoreGet(w, r, table)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay toggle score handler for GET method.
func (server *Server) handleOverlayToggleScoreGet(w http.ResponseWriter, r *http.Request, table int) {
	showScore := server.tables[table].Overlay.ToggleScore()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[table].Overlay,
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	server.handleSuccess(w, r, &OverlayToggleScoreResp{
		ShowScore: showScore,
	})
}
