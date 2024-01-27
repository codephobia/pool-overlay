package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/libs/go/challonge"
	"github.com/codephobia/pool-overlay/libs/go/models"
	"github.com/gorilla/mux"
)

type TournamentLoadBody struct {
	ID       int                        `json:"id"`
	Settings TournamentLoadBodySettings `json:"settings"`
}

type TournamentLoadBodySettings struct {
	MaxTables     int             `json:"max_tables"`
	GameType      models.GameType `json:"game_type"`
	ShowOverlay   bool            `json:"show_overlay"`
	ShowFlags     bool            `json:"show_flags"`
	ShowFargo     bool            `json:"show_fargo"`
	ShowScore     bool            `json:"show_score"`
	IsHandicapped bool            `json:"is_handicapped"`
	ASideRaceTo   int             `json:"a_side_race_to"`
	BSideRaceTo   int             `json:"b_side_race_to"`
}

// Handler for /tournament.
func (server *Server) handleTournament() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleTournamentGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Tournament handler for GET method.
func (server *Server) handleTournamentGet(w http.ResponseWriter, r *http.Request) {
	// check for loaded tournament
	if server.challonge.Tournament == nil {
		server.handleError(w, r, http.StatusNotFound, ErrTournamentNotFound)
		return
	}

	// send response
	server.handleSuccess(w, r, server.challonge.Tournament)
}

// Handler for /tournament/list.
func (server *Server) handleTournamentList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleTournamentListGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Tournaments list handler for GET method.
func (server *Server) handleTournamentListGet(w http.ResponseWriter, r *http.Request) {
	// get unfinished tournaments from Challonge API
	tournaments, err := server.challonge.GetTournamentList()
	if err != nil {
		server.handleError(w, r, http.StatusNotFound, err)
		return
	}

	// send response
	server.handleSuccess(w, r, tournaments)
}

// Handler for /tournament/{tournamentID}.
func (server *Server) handleTournamentByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleTournamentByIDGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// TournamentByID handler for GET method.
func (server *Server) handleTournamentByIDGet(w http.ResponseWriter, r *http.Request) {
	// get param for tournament id from url
	params := mux.Vars(r)
	tournamentID, ok := params["tournamentID"]
	if !ok || len(tournamentID) == 0 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTournamentID)
		return
	}

	// convert tournament id to int
	id, err := strconv.Atoi(tournamentID)
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTournamentID)
		return
	}

	// get tournament by id
	tournament, err := server.challonge.GetTournamentByID(id)
	if err != nil {
		server.handleError(w, r, http.StatusNotFound, err)
		return
	}

	// send response
	server.handleSuccess(w, r, tournament)
}

// Handler for /tournament/load.
func (server *Server) handleTournamentLoad() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "POST":
			server.handleTournamentLoadPost(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Tournament load handler for POST method.
func (server *Server) handleTournamentLoadPost(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body TournamentLoadBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidTournamentDetails)
		return
	}

	// make settings
	settings := &challonge.Settings{
		MaxTables:     body.Settings.MaxTables,
		GameType:      body.Settings.GameType,
		ShowOverlay:   body.Settings.ShowOverlay,
		ShowFlags:     body.Settings.ShowFlags,
		ShowFargo:     body.Settings.ShowFargo,
		ShowScore:     body.Settings.ShowScore,
		IsHandicapped: body.Settings.IsHandicapped,
		ASideRaceTo:   body.Settings.ASideRaceTo,
		BSideRaceTo:   body.Settings.BSideRaceTo,
	}

	// load tournament
	err := server.challonge.LoadTournament(body.ID, settings)
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	// send response
	server.handle204Success(w, r)
}

// Handler for /tournament/unload.
func (server *Server) handleTournamentUnload() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "POST":
			server.handleTournamentUnloadPost(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Tournament unload handler for POST method.
func (server *Server) handleTournamentUnloadPost(w http.ResponseWriter, r *http.Request) {
	// unload the tournament
	server.challonge.UnloadTournament()

	// send response
	server.handle204Success(w, r)
}
