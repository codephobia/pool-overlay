package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/codephobia/pool-overlay/libs/go/events"
	"github.com/codephobia/pool-overlay/libs/go/models"
	"github.com/codephobia/pool-overlay/libs/go/overlay"
)

const (
	gameDirectionIncrement = "increment"
	gameDirectionDecrement = "decrement"
)

// GameTypePatchBody is the incoming body on a patch request for updating the
// game type.
type GameTypePatchBody struct {
	Type models.GameType `json:"type"`
}

// GameTypeResp is a reponse containing the game type.
type GameTypeResp struct {
	Type models.GameType `json:"type"`
}

// GameVsModePatchBody is the incoming body on a patch request for updating the
// game vs mode.
type GameVsModePatchBody struct {
	VsMode models.GameVsMode `json:"vsMode"`
}

// GameVsModeResp is a reponse containing the game type.
type GameVsModeResp struct {
	VsMode models.GameVsMode `json:"vsMode"`
}

// GameRaceToPatchBody is the incoming body on a patch request for updating the
// game race to.
type GameRaceToPatchBody struct {
	Direction string `json:"direction"`
}

// GameRaceToResp is a reponse containing the game race to.
type GameRaceToResp struct {
	RaceTo int `json:"raceTo"`
}

// GameScorePatchBody is the incoming body on a patch request for updating the
// game score for the specified player.
type GameScorePatchBody struct {
	PlayerNum int    `json:"playerNum"`
	Direction string `json:"direction"`
}

// GameScoreResp is a reponse containing the game score.
type GameScoreResp struct {
	ScoreOne int `json:"scoreOne"`
	ScoreTwo int `json:"scoreTwo"`
}

// GamePlayersPatchBody is the incoming body on a patch request for updating the
// player num to a specified player id.
type GamePlayersPatchBody struct {
	PlayerNum int `json:"playerNum"`
	PlayerID  int `json:"playerID"`
}

// GamePlayersDeleteBody is the incoming body on a delete request for unsetting
// the specified player num.
type GamePlayersDeleteBody struct {
	PlayerNum int `json:"playerNum"`
}

// GamePlayersFlagPatchBody is the incoming body on a patch request for updating
// the player num to a specified flag id.
type GamePlayersFlagPatchBody struct {
	PlayerNum int `json:"playerNum"`
	FlagID    int `json:"flagID"`
}

// GamePlayersNamePatchBody is the incoming body on a patch request for updating
// the player name for the specified player num.
type GamePlayersNamePatchBody struct {
	PlayerNum int    `json:"playerNum"`
	Name      string `json:"name"`
}

type GameTeamsPatchBody struct {
	TeamNum int `json:"teamNum"`
	TeamID  int `json:"teamID"`
}

// Handler for /game.
func (server *Server) handleGame() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleGameGet(w, r)
		case "POST":
			server.handleGamePost(w, r)
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

// Game handler for POST method.
func (server *Server) handleGamePost(w http.ResponseWriter, r *http.Request) {
	// Save existing game.
	if err := server.state.Game.Save(true); err != nil {
		server.handleError(w, r, http.StatusInternalServerError, err)
		return
	}

	// Reset game to create a new one with same players / settings.
	if err := server.state.Game.Reset(); err != nil {
		server.handleError(w, r, http.StatusInternalServerError, err)
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
	server.handle204Success(w, r)
}

// Handler for /game/type.
func (server *Server) handleGameType() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGameTypePatch(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game type handler for PATCH method.
func (server *Server) handleGameTypePatch(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GameTypePatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// check that game type is valid
	if !body.Type.IsValid() {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// update game type
	if err := server.state.Game.SetType(body.Type); err != nil {
		// TODO: LOG THIS ERROR
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
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
	server.handleSuccess(w, r, GameTypeResp(body))
}

// Handler for /game/vs-mode.
func (server *Server) handleGameVsMode() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGameVsModePatch(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game vs-mode handler for PATCH method.
func (server *Server) handleGameVsModePatch(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GameVsModePatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// check the vs mode is valid
	if !body.VsMode.IsValid() {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameVsMode)
		return
	}

	// update game vs mode
	if err := server.state.Game.SetVsMode(body.VsMode); err != nil {
		// TODO: LOG THIS ERROR
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
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
	server.handleSuccess(w, r, GameVsModeResp(body))
}

// Handler for /game/race-to.
func (server *Server) handleGameRaceTo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGameRaceToPatch(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game race-to handler for PATCH method.
func (server *Server) handleGameRaceToPatch(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GameRaceToPatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// validate direction
	if body.Direction != gameDirectionIncrement && body.Direction != gameDirectionDecrement {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameDirection)
		return
	}

	// update race to number
	if body.Direction == gameDirectionIncrement {
		if err := server.state.Game.IncrementRaceTo(); err != nil {
			// TODO: LOG THIS ERROR
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		}
	} else {
		if err := server.state.Game.DecrementRaceTo(); err != nil {
			// TODO: LOG THIS ERROR
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
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
	server.handleSuccess(w, r, GameRaceToResp{
		RaceTo: server.state.Game.RaceTo,
	})
}

// Handler for /game/score
func (server *Server) handleGameScore() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGameScorePatch(w, r)
		case "DELETE":
			server.handleGameScoreDelete(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game score handler for PATCH method.
func (server *Server) handleGameScorePatch(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GameScorePatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// validate direction
	if body.Direction != gameDirectionIncrement && body.Direction != gameDirectionDecrement {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameDirection)
		return
	}

	// update score
	if body.Direction == gameDirectionIncrement {
		if err := server.state.Game.IncrementScore(body.PlayerNum); err != nil {
			if errors.Is(err, models.ErrInvalidPlayerNumber) {
				server.handleError(w, r, http.StatusUnprocessableEntity, err)
			} else {
				// TODO: LOG THIS ERROR
				server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			}
			return
		}
	} else {
		if err := server.state.Game.DecrementScore(body.PlayerNum); err != nil {
			if errors.Is(err, models.ErrInvalidPlayerNumber) {
				server.handleError(w, r, http.StatusUnprocessableEntity, err)
			} else {
				// TODO: LOG THIS ERROR
				server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			}
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

// Game score reset handler for DELETE method.
func (server *Server) handleGameScoreDelete(w http.ResponseWriter, r *http.Request) {
	// reset game score
	if err := server.state.Game.ResetScore(); err != nil {
		// TODO: LOG THIS ERROR
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
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

// Handler for /game/players
func (server *Server) handleGamePlayers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGamePlayersPatch(w, r)
		case "DELETE":
			server.handleGamePlayersDelete(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game players handler for PATCH method.
func (server *Server) handleGamePlayersPatch(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GamePlayersPatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	var player models.Player
	if err := player.LoadByID(server.db, body.PlayerID); err != nil {
		// TODO: MAYBE CHANGE THIS TO ERRORS.IS
		if err == models.ErrPlayerIDInvalid {
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			return
		}
		if err == models.ErrPlayerNotFound {
			server.handleError(w, r, http.StatusNotFound, ErrPlayerNotFound)
			return
		}
	}

	if err := server.state.Game.SetPlayer(body.PlayerNum, &player); err != nil {
		if errors.Is(err, models.ErrInvalidPlayerNumber) {
			server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerNumber)
		} else {
			// TODO: LOG THIS ERROR
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		}
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

// Game players handler for DELETE method.
func (server *Server) handleGamePlayersDelete(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GamePlayersDeleteBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// unset the current player
	if err := server.state.Game.UnsetPlayer(body.PlayerNum); err != nil {
		if errors.Is(err, models.ErrInvalidPlayerNumber) {
			server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerNumber)
		} else {
			// TODO: LOG THIS ERROR
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		}
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

// Handler for /game/players/flag
func (server *Server) handleGamePlayersFlag() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGamePlayersFlagPatch(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game players flag handler for PATCH method.
func (server *Server) handleGamePlayersFlagPatch(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GamePlayersFlagPatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	var flag models.Flag
	if err := flag.LoadByID(server.db, body.FlagID); err != nil {
		if err == models.ErrFlagIDInvalid {
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			return
		}
		if err == models.ErrFlagNotFound {
			server.handleError(w, r, http.StatusNotFound, ErrPlayerNotFound)
			return
		}
	}

	if err := server.state.Game.SetPlayerFlag(body.PlayerNum, &flag); err != nil {
		if errors.Is(err, models.ErrInvalidPlayerNumber) {
			server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidPlayerNumber)
		} else {
			// TODO: LOG THIS ERROR
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		}
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

// Handler for /game/players/name
func (server *Server) handleGamePlayersName() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGamePlayersNamePatch(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game players name handler for PATCH method.
func (server *Server) handleGamePlayersNamePatch(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GamePlayersNamePatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	if err := server.state.Game.SetPlayerName(body.PlayerNum, body.Name); err != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
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

// Handler for /game/teams
func (server *Server) handleGameTeams() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGameTeamsPatch(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game teams handler for PATCH method.
func (server *Server) handleGameTeamsPatch(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body GameTeamsPatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	var team models.Team
	if err := team.LoadByID(server.db, body.TeamID); err != nil {
		if err == models.ErrTeamIDInvalid {
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			return
		}
		if err == models.ErrTeamNotFound {
			server.handleError(w, r, http.StatusNotFound, ErrTeamNotFound)
			return
		}
	}

	if err := server.state.Game.SetTeam(body.TeamNum, &team); err != nil {
		if errors.Is(err, models.ErrInvalidTeamNumber) {
			server.handleError(w, r, http.StatusUnprocessableEntity, models.ErrInvalidTeamNumber)
		} else {
			// TODO: LOG THIS ERROR
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		}
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
