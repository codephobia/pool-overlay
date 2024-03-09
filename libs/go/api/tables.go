package api

import (
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/libs/go/events"
	"github.com/codephobia/pool-overlay/libs/go/overlay"
	"github.com/codephobia/pool-overlay/libs/go/state"
	"github.com/gorilla/mux"
)

// Handler for /table/count.
func (server *Server) handleTableCount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleTableCountGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Table count handler for GET method.
func (server *Server) handleTableCountGet(w http.ResponseWriter, r *http.Request) {
	count := int64(len(server.tables))
	server.handleSuccess(w, r, &CountResp{
		Count: count,
	})
}

// Handler for /table/add.
func (server *Server) handleTableAdd() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "POST":
			server.handleTableAddPost(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Table add handler for POST method.
func (server *Server) handleTableAddPost(w http.ResponseWriter, r *http.Request) {
	count := len(server.tables)
	server.tables[count+1] = state.NewState(server.db, count+1)

	server.handleSuccess(w, r, &CountResp{
		Count: int64(len(server.tables)),
	})
}

// Handler for /table/remove.
func (server *Server) handleTableRemove() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "POST":
			server.handleTableRemovePost(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Table remove handler for POST method.
func (server *Server) handleTableRemovePost(w http.ResponseWriter, r *http.Request) {
	count := len(server.tables)

	// don't allow removal of single table
	if count == 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrRemoveOnlyTable)
		return
	}

	delete(server.tables, count)

	server.handleSuccess(w, r, &CountResp{
		Count: int64(len(server.tables)),
	})
}

// Handler for /table/swap.
func (server *Server) handleTableSwap() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleTableSwapGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Table swap handler for GET method.
func (server *Server) handleTableSwapGet(w http.ResponseWriter, r *http.Request) {
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

	// get new table number
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
	server.tables[tableNum].Table = newTable
	server.tables[tableNum].Game.Table = newTable
	server.tables[tableNum].Overlay.Table = newTable

	server.tables[newTable].Table = tableNum
	server.tables[newTable].Game.Table = tableNum
	server.tables[newTable].Overlay.Table = tableNum

	server.tables[tableNum], server.tables[newTable] = server.tables[newTable], server.tables[tableNum]

	// update current matches in tournament mode if running
	if server.challonge.InTournamentMode() {
		server.challonge.CurrentMatches[tableNum], server.challonge.CurrentMatches[newTable] = server.challonge.CurrentMatches[newTable], server.challonge.CurrentMatches[tableNum]
	}

	// broadcast out changes for each table
	changedTables := []int{tableNum, newTable}

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
