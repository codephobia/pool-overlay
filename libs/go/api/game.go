package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/libs/go/events"
	"github.com/codephobia/pool-overlay/libs/go/models"
	"github.com/codephobia/pool-overlay/libs/go/overlay"
	"github.com/gorilla/mux"
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
	RaceTo              int  `json:"raceTo"`
	UseFargoHotHandicap bool `json:"useFargoHotHandicap"`
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

// GameTeamsPatchBody is the incoming body on a patch request for updating the
// team for the specified team number.
type GameTeamsPatchBody struct {
	TeamNum int `json:"teamNum"`
	TeamID  int `json:"teamID"`
}

// GameFargoHotHandicapPatchBody is the incoming body on a patch request for
// updating the fargo hot handicap.
type GameFargoHotHandicapPatchBody struct {
	UseFargoHotHandicap bool `json:"useFargoHotHandicap"`
}

// GameFargoHotHandicapPatchResp is the response containing the new fargo hot
// handicap setting.
type GameFargoHotHandicapPatchResp struct {
	UseFargoHotHandicap bool `json:"useFargoHotHandicap"`
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

	// send response
	server.handleSuccess(w, r, server.tables[tableNum].Game)
}

// Game handler for POST method.
func (server *Server) handleGamePost(w http.ResponseWriter, r *http.Request) {
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

	// Save existing game.
	if err := server.tables[tableNum].Game.Save(true); err != nil {
		server.handleError(w, r, http.StatusInternalServerError, err)
		return
	}

	// Check if we're in tournament mode right now.
	if server.challonge.InTournamentMode() {
		if err := server.challonge.Continue(tableNum); err != nil {
			log.Printf("unable to continue tournament: %s", err)
		}
	} else {
		// Reset game to create a new one with same players / settings.
		if err := server.tables[tableNum].Game.Reset(); err != nil {
			server.handleError(w, r, http.StatusInternalServerError, err)
			return
		}
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.tables[tableNum].Game),
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
	if err := server.tables[tableNum].Game.SetType(body.Type); err != nil {
		// TODO: LOG THIS ERROR
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.tables[tableNum].Game),
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
	if err := server.tables[tableNum].Game.SetVsMode(body.VsMode); err != nil {
		// TODO: LOG THIS ERROR
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.tables[tableNum].Game),
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
		if err := server.tables[tableNum].Game.IncrementRaceTo(); err != nil {
			// TODO: LOG THIS ERROR
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		}
	} else {
		if err := server.tables[tableNum].Game.DecrementRaceTo(); err != nil {
			// TODO: LOG THIS ERROR
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		}
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.tables[tableNum].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, GameRaceToResp{
		RaceTo:              server.tables[tableNum].Game.RaceTo,
		UseFargoHotHandicap: server.tables[tableNum].Game.UseFargoHotHandicap,
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
		if err := server.tables[tableNum].Game.IncrementScore(body.PlayerNum); err != nil {
			if errors.Is(err, models.ErrInvalidPlayerNumber) {
				server.handleError(w, r, http.StatusUnprocessableEntity, err)
			} else {
				// TODO: LOG THIS ERROR
				server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			}
			return
		}
	} else {
		if err := server.tables[tableNum].Game.DecrementScore(body.PlayerNum); err != nil {
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
		events.NewGameEventPayload(server.tables[tableNum].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// Check if we're in tournament mode right now.
	if server.challonge.InTournamentMode() {
		if err := server.challonge.UpdateMatchScore(tableNum); err != nil {
			// fail gracefully since live score keeping isn't that important
			log.Printf("error updating match score on challonge: %s", err)
		}
	}

	// send response
	server.handleSuccess(w, r, GameScoreResp{
		ScoreOne: server.tables[tableNum].Game.ScoreOne,
		ScoreTwo: server.tables[tableNum].Game.ScoreTwo,
	})
}

// Game score reset handler for DELETE method.
func (server *Server) handleGameScoreDelete(w http.ResponseWriter, r *http.Request) {
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

	// reset game score
	if err := server.tables[tableNum].Game.ResetScore(); err != nil {
		// TODO: LOG THIS ERROR
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.tables[tableNum].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, GameScoreResp{
		ScoreOne: server.tables[tableNum].Game.ScoreOne,
		ScoreTwo: server.tables[tableNum].Game.ScoreTwo,
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

	if err := server.tables[tableNum].Game.SetPlayer(body.PlayerNum, &player); err != nil {
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
		events.NewGameEventPayload(server.tables[tableNum].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, server.tables[tableNum].Game)
}

// Game players handler for DELETE method.
func (server *Server) handleGamePlayersDelete(w http.ResponseWriter, r *http.Request) {
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

	// decode the body
	var body GamePlayersDeleteBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// unset the current player
	if err := server.tables[tableNum].Game.UnsetPlayer(body.PlayerNum); err != nil {
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
		events.NewGameEventPayload(server.tables[tableNum].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, server.tables[tableNum].Game)
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

	if err := server.tables[1].Game.SetPlayerFlag(body.PlayerNum, &flag); err != nil {
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
		events.NewGameEventPayload(server.tables[1].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, server.tables[1].Game)
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

	if err := server.tables[1].Game.SetPlayerName(body.PlayerNum, body.Name); err != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.tables[1].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, server.tables[1].Game)
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

	if err := server.tables[1].Game.SetTeam(body.TeamNum, &team); err != nil {
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
		events.NewGameEventPayload(server.tables[1].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, server.tables[1].Game)
}

// Handler for /game/fargo-hot-handicap.
func (server *Server) handleGameFargoHotHandicap() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "PATCH":
			server.handleGameFargoHotHandicapPatch(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Game fargo-hot-handicap handler for PATCH method.
func (server *Server) handleGameFargoHotHandicapPatch(w http.ResponseWriter, r *http.Request) {
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

	// decode the body
	var body GameFargoHotHandicapPatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidGameType)
		return
	}

	// Update the game fargo hot handicap option.
	if err := server.tables[tableNum].Game.SetUseFargoHotHandicap(body.UseFargoHotHandicap); err != nil {
		// TODO: Currently all errors return as 500 here, but might not always make sense. Could use errors.Is for this.
		server.handleError(w, r, http.StatusInternalServerError, err)
		return
	}

	// Generate message to broadcast to overlay.
	message, err := overlay.NewEvent(
		events.GameEventType,
		events.NewGameEventPayload(server.tables[tableNum].Game),
	).ToBytes()
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrUnableToBroadcastUpdate)
		return
	}

	// Broadcast update to overlay.
	server.overlay.Broadcast <- message

	// send response
	server.handleSuccess(w, r, GameFargoHotHandicapPatchResp{
		UseFargoHotHandicap: server.tables[tableNum].Game.UseFargoHotHandicap,
	})
}
