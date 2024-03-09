package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/libs/go/events"
	"github.com/codephobia/pool-overlay/libs/go/overlay"
	"github.com/gorilla/mux"
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

	// On connect, send all table games and overlay states.
	for i := 1; i <= len(server.tables); i++ {
		// Generate current game state message.
		gameMessage, err := overlay.NewEvent(
			events.GameEventType,
			events.NewGameEventPayload(server.tables[i].Game),
		).ToBytes()
		if err != nil {
			server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
			return
		}

		// Send current state to new connection.
		oc.Send <- gameMessage

		// Generate current overlay state message.
		message, err := overlay.NewEvent(
			events.OverlayStateEventType,
			server.tables[i].Overlay,
		).ToBytes()
		if err != nil {
			server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
			return
		}

		// Send current state to new connection.
		oc.Send <- message
	}
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
	// get params for table numbers from url
	params := mux.Vars(r)

	// get table number
	tableNumValue, ok := params["tableNum"]
	if !ok || len(tableNumValue) < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	// convert table number to int
	tableNum, err := strconv.Atoi(tableNumValue)
	if err != nil || tableNum > len(server.tables) {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	hidden := server.tables[tableNum].Overlay.ToggleHidden()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[tableNum].Overlay,
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
func (server *Server) handleOverlayToggleFlags() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayToggleFlagsGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay toggle flags handler for GET method.
func (server *Server) handleOverlayToggleFlagsGet(w http.ResponseWriter, r *http.Request) {
	// get params for table numbers from url
	params := mux.Vars(r)

	// get table number
	tableNumValue, ok := params["tableNum"]
	if !ok || len(tableNumValue) < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	// convert table number to int
	tableNum, err := strconv.Atoi(tableNumValue)
	if err != nil || tableNum > len(server.tables) {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	showFlags := server.tables[tableNum].Overlay.ToggleFlags()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[tableNum].Overlay,
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
func (server *Server) handleOverlayToggleFargo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayToggleFargoGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay toggle fargo handler for GET method.
func (server *Server) handleOverlayToggleFargoGet(w http.ResponseWriter, r *http.Request) {
	// get params for table numbers from url
	params := mux.Vars(r)

	// get table number
	tableNumValue, ok := params["tableNum"]
	if !ok || len(tableNumValue) < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	// convert table number to int
	tableNum, err := strconv.Atoi(tableNumValue)
	if err != nil || tableNum > len(server.tables) {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	showFargo := server.tables[tableNum].Overlay.ToggleFargo()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[tableNum].Overlay,
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
func (server *Server) handleOverlayToggleScore() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleOverlayToggleScoreGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Overlay toggle score handler for GET method.
func (server *Server) handleOverlayToggleScoreGet(w http.ResponseWriter, r *http.Request) {
	// get params for table numbers from url
	params := mux.Vars(r)

	// get table number
	tableNumValue, ok := params["tableNum"]
	if !ok || len(tableNumValue) < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	// convert table number to int
	tableNum, err := strconv.Atoi(tableNumValue)
	if err != nil || tableNum > len(server.tables) {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	showScore := server.tables[tableNum].Overlay.ToggleScore()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.OverlayStateEventType,
		server.tables[tableNum].Overlay,
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
