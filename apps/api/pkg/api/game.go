package api

import (
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/apps/api/pkg/game"
	"github.com/codephobia/pool-overlay/apps/api/pkg/models"
)

const (
	gameDirectionIncrement = "increment"
	gameDirectionDecrement = "decrement"
)

// GameTypeResp is a reponse containing the game type.
type GameTypeResp struct {
	Type game.Type `json:"type"`
}

// GameVsModeResp is a reponse containing the game type.
type GameVsModeResp struct {
	VsMode game.VsMode `json:"vs_mode"`
}

// GameRaceToResp is a reponse containing the game race to.
type GameRaceToResp struct {
	RaceTo int `json:"race_to"`
}

// GameScoreResp is a reponse containing the game score.
type GameScoreResp struct {
	Score game.Score `json:"score"`
}

// Handler for /game.
func (server *Server) handleGame() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleGameGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game handler for GET method.
func (server *Server) handleGameGet(w http.ResponseWriter, r *http.Request) {
	// send response
	server.handleSuccess(w, r, server.game)
}

// Handler for /game/update/type.
func (server *Server) handleGameType() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleGameTypeGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game type handler for GET method.
func (server *Server) handleGameTypeGet(w http.ResponseWriter, r *http.Request) {
	// get query vars
	v := r.URL.Query()

	// get game type from query params
	gameType, err := strconv.ParseUint(v.Get("type"), 10, 0)
	if err != nil || !game.Type(gameType).IsValid() {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// update game type
	server.game.SetType(game.Type(gameType))

	// TODO: BROADCAST OVERLAY UPDATE

	// send response
	server.handleSuccess(w, r, GameTypeResp{
		Type: game.Type(gameType),
	})
}

// Handler for /game/update/vs-mode.
func (server *Server) handleGameVsMode() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleGameVsModeGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game vs-mode handler for GET method.
func (server *Server) handleGameVsModeGet(w http.ResponseWriter, r *http.Request) {
	// get query vars
	v := r.URL.Query()

	// get game mode from query params
	gameMode, err := strconv.ParseUint(v.Get("mode"), 10, 0)
	if err != nil || !game.VsMode(gameMode).IsValid() {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameVsMode)
		return
	}

	// update game mode
	server.game.SetVsMode(game.VsMode(gameMode))

	// TODO: BROADCAST OVERLAY UPDATE

	// send response
	server.handleSuccess(w, r, GameVsModeResp{
		VsMode: game.VsMode(gameMode),
	})
}

// Handler for /game/update/race-to.
func (server *Server) handleGameRaceTo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleGameRaceToGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game race-to handler for GET method.
func (server *Server) handleGameRaceToGet(w http.ResponseWriter, r *http.Request) {
	// get query vars
	v := r.URL.Query()

	// get game race to direction from query params
	direction := v.Get("direction")

	// valid directions
	validDirections := make(map[string]struct{})
	validDirections[gameDirectionIncrement] = struct{}{}
	validDirections[gameDirectionDecrement] = struct{}{}

	// validate direction
	if _, ok := validDirections[direction]; !ok {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameDirection)
		return
	}

	// update race to number
	if direction == gameDirectionIncrement {
		server.game.IncrementRaceTo()
	} else {
		server.game.DecrementRaceTo()
	}

	// TODO: BROADCAST OVERLAY UPDATE

	// send response
	server.handleSuccess(w, r, GameRaceToResp{
		RaceTo: server.game.RaceTo,
	})
}

// Handler for /game/update/score
func (server *Server) handleGameScore() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleGameScoreGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game score handler for GET method.
func (server *Server) handleGameScoreGet(w http.ResponseWriter, r *http.Request) {
	// get query vars
	v := r.URL.Query()

	// get player num to from query params
	playerNum, err := strconv.Atoi(v.Get("playerNum"))
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, game.ErrInvalidPlayerNumber)
		return
	}

	// get game race to from query params
	direction := v.Get("direction")

	// valid directions
	validDirections := make(map[string]struct{})
	validDirections[gameDirectionIncrement] = struct{}{}
	validDirections[gameDirectionDecrement] = struct{}{}

	// validate direction
	if _, ok := validDirections[direction]; !ok {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameDirection)
		return
	}

	// update score
	if direction == gameDirectionIncrement {
		if err := server.game.IncrementScore(playerNum); err != nil {
			server.handleError(w, r, http.StatusUnprocessableEntity, err)
			return
		}
	} else {
		if err := server.game.DecrementScore(playerNum); err != nil {
			server.handleError(w, r, http.StatusUnprocessableEntity, err)
			return
		}
	}

	// TODO: BROADCAST OVERLAY UPDATE

	// send response
	server.handleSuccess(w, r, GameScoreResp{
		Score: server.game.Score,
	})
}

// Handler for /game/update/score/reset
func (server *Server) handleGameScoreReset() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleGameScoreResetGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game score reset handler for GET method.
func (server *Server) handleGameScoreResetGet(w http.ResponseWriter, r *http.Request) {
	// reset game score
	server.game.ResetScore()

	// TODO: BROADCAST OVERLAY UPDATE

	// send response
	server.handleSuccess(w, r, GameScoreResp{
		Score: server.game.Score,
	})
}

// Handler for /game/update/player
func (server *Server) handleGamePlayers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleGamePlayersGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game player handler for GET method.
func (server *Server) handleGamePlayersGet(w http.ResponseWriter, r *http.Request) {
	// get query vars
	v := r.URL.Query()

	// get player num from query params
	playerNum, err := strconv.Atoi(v.Get("playerNum"))
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, game.ErrInvalidPlayerNumber)
		return
	}

	// get player id from query params
	playerID, err := strconv.Atoi(v.Get("playerID"))
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, game.ErrInvalidPlayerID)
		return
	}

	var player models.Player
	if err := player.LoadByID(server.db, playerID); err != nil {
		if err == models.ErrPlayerIDInvalid {
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			return
		}
		if err == models.ErrPlayerNotFound {
			server.handleError(w, r, http.StatusNotFound, ErrPlayerNotFound)
			return
		}
	}

	if err := server.game.SetPlayer(playerNum, &player); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, game.ErrInvalidPlayerNumber)
		return
	}

	// TODO: BROADCAST OVERLAY UPDATE

	// send response
	server.handleSuccess(w, r, server.game)
}

// Handler for /game/update/teams
func (server *Server) handleGameTeams() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handleGameTeamsGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game teams handler for GET method.
func (server *Server) handleGameTeamsGet(w http.ResponseWriter, r *http.Request) {
	// get query vars
	v := r.URL.Query()

	// get team num from query params
	teamNum, err := strconv.Atoi(v.Get("teamNum"))
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, game.ErrInvalidPlayerNumber)
		return
	}

	// get team id from query params
	teamID, err := strconv.Atoi(v.Get("teamID"))
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, game.ErrInvalidPlayerID)
		return
	}

	var team models.Team
	if err := team.LoadByID(server.db, teamID); err != nil {
		if err == models.ErrTeamIDInvalid {
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			return
		}
		if err == models.ErrTeamNotFound {
			server.handleError(w, r, http.StatusNotFound, ErrTeamNotFound)
			return
		}
	}

	if err := server.game.SetTeam(teamNum, &team); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, game.ErrInvalidTeamNumber)
		return
	}

	// TODO: BROADCAST OVERLAY UPDATE

	// send response
	server.handleSuccess(w, r, server.game)
}
