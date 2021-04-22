package api

import (
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/apps/api/pkg/events"
	"github.com/codephobia/pool-overlay/apps/api/pkg/models"
	"github.com/codephobia/pool-overlay/apps/api/pkg/overlay"
)

const (
	gameDirectionIncrement = "increment"
	gameDirectionDecrement = "decrement"
)

// GameTypeResp is a reponse containing the game type.
type GameTypeResp struct {
	Type models.GameType `json:"type"`
}

// GameVsModeResp is a reponse containing the game type.
type GameVsModeResp struct {
	VsMode models.GameVsMode `json:"vs_mode"`
}

// GameRaceToResp is a reponse containing the game race to.
type GameRaceToResp struct {
	RaceTo int `json:"race_to"`
}

// GameScoreResp is a reponse containing the game score.
type GameScoreResp struct {
	ScoreOne int `json:"score_one"`
	ScoreTwo int `json:"score_two"`
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
	server.handleSuccess(w, r, server.state.Game)
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
	if err != nil || !models.GameType(gameType).IsValid() {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// update game type
	server.state.Game.SetType(models.GameType(gameType))

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.state.Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, GameTypeResp{
		Type: models.GameType(gameType),
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
	if err != nil || !models.GameVsMode(gameMode).IsValid() {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameVsMode)
		return
	}

	// update game mode
	server.state.Game.SetVsMode(models.GameVsMode(gameMode))

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.state.Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, GameVsModeResp{
		VsMode: models.GameVsMode(gameMode),
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
		server.state.Game.IncrementRaceTo()
	} else {
		server.state.Game.DecrementRaceTo()
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.state.Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, GameRaceToResp{
		RaceTo: server.state.Game.RaceTo,
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
		server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerNumber)
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
		if err := server.state.Game.IncrementScore(playerNum); err != nil {
			server.handleError(w, r, http.StatusUnprocessableEntity, err)
			return
		}
	} else {
		if err := server.state.Game.DecrementScore(playerNum); err != nil {
			server.handleError(w, r, http.StatusUnprocessableEntity, err)
			return
		}
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.state.Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, GameScoreResp{
		ScoreOne: server.state.Game.ScoreOne,
		ScoreTwo: server.state.Game.ScoreTwo,
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
	server.state.Game.ResetScore()

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.state.Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, GameScoreResp{
		ScoreOne: server.state.Game.ScoreOne,
		ScoreTwo: server.state.Game.ScoreTwo,
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
		server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerNumber)
		return
	}

	// get player id from query params
	playerID, err := strconv.Atoi(v.Get("playerID"))
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerID)
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

	if err := server.state.Game.SetPlayer(playerNum, &player); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerNumber)
		return
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.state.Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, server.state.Game)
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
		server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerNumber)
		return
	}

	// get team id from query params
	teamID, err := strconv.Atoi(v.Get("teamID"))
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerID)
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

	if err := server.state.Game.SetTeam(teamNum, &team); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidTeamNumber)
		return
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.state.Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, server.state.Game)
}
