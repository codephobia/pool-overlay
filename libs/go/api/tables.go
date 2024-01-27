package api

import (
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/libs/go/events"
	"github.com/codephobia/pool-overlay/libs/go/overlay"
	"github.com/gorilla/mux"
)

// Handler for /table/swap.
func (server *Server) handleTableSwap(table int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleTableSwapGet(w, r, table)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Table swap handler for GET method.
func (server *Server) handleTableSwapGet(w http.ResponseWriter, r *http.Request, table int) {
	// get param for new table num from url
	params := mux.Vars(r)
	newTableValue, ok := params["newTable"]
	if !ok || len(newTableValue) < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	// convert new table to int
	newTable, err := strconv.Atoi(newTableValue)
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTableNumber)
		return
	}

	// swap table states
	server.tables[table].Table = newTable
	server.tables[table].Game.Table = newTable
	server.tables[table].Overlay.Table = newTable

	server.tables[newTable].Table = table
	server.tables[newTable].Game.Table = table
	server.tables[newTable].Overlay.Table = table

	server.tables[table], server.tables[newTable] = server.tables[newTable], server.tables[table]

	// update current matches in tournament mode if running
	if server.challonge.InTournamentMode() {
		server.challonge.CurrentMatches[table], server.challonge.CurrentMatches[newTable] = server.challonge.CurrentMatches[newTable], server.challonge.CurrentMatches[table]
	}

	// broadcast out changes for each table
	changedTables := []int{table, newTable}

	for i := 0; i < len(changedTables); i++ {
		// Generate current game state message.
		gameMessage, err := overlay.NewEvent(
			events.GameEventType,
			events.NewGameEventPayload(server.tables[changedTables[i]].Game),
		).ToBytes()
		if err != nil {
			server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
			return
		}

		// Send current state to table.
		server.overlay.Broadcast <- gameMessage

		// Generate current overlay state message.
		message, err := overlay.NewEvent(
			events.OverlayStateEventType,
			server.tables[changedTables[i]].Overlay,
		).ToBytes()
		if err != nil {
			server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
			return
		}

		// Send current state to table.
		server.overlay.Broadcast <- message
	}

	// return results
	server.handle204Success(w, r)
}
