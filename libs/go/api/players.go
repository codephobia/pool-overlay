package api

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/libs/go/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const (
	playersPerPage = 100
)

// PlayersPostBody is an incoming body on a POST request for creating a player.
type PlayersPostBody struct {
	Name        string `json:"name"`
	FlagID      uint   `json:"flag_id"`
	FargoID     uint   `json:"fargo_id"`
	FargoRating uint   `json:"fargo_rating"`
}

// PlayersPatchBody is an incoming body on a PATCH request for updating a
// player.
type PlayersPatchBody struct {
	Name        string `json:"name"`
	FlagID      uint   `json:"flag_id"`
	FargoID     uint   `json:"fargo_id"`
	FargoRating uint   `json:"fargo_rating"`
}

// Handler for /players.
func (server *Server) handlePlayers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handlePlayersGet(w, r)
		case "POST":
			server.handlePlayersPost(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Players handler for GET method. Returns a page of players.
func (server *Server) handlePlayersGet(w http.ResponseWriter, r *http.Request) {
	// get query vars
	v := r.URL.Query()

	// get page
	page := v.Get("page")

	// default page to 1 if it doesn't exist
	if page == "" {
		page = "1"
	}

	// convert page number and validate
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPageNumber)
		return
	}

	// get count of players to test page ceiling
	var count int64
	countResult := server.db.Model(&models.Player{}).Count(&count)
	if countResult.Error != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	// check if page is beyond maximum
	totalPages := int(math.Ceil(float64(count) / playersPerPage))
	if pageNum > totalPages {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPageNumber)
		return
	}

	// get players for page
	var players []*models.Player

	offset := pageNum*playersPerPage - playersPerPage
	playersResult := server.db.
		Select("id", "name", "flag_id", "fargo_id", "fargo_rating").
		Order("name").
		Limit(playersPerPage).
		Offset(offset).
		Preload("Flag", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "country", "image_path")
		}).
		Find(&players)
	if playersResult.Error != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	// return results
	server.handleSuccess(w, r, players)
}

// Players handler for POST method. Creates a new player.
func (server *Server) handlePlayersPost(w http.ResponseWriter, r *http.Request) {
	// decode the body
	var body PlayersPostBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPlayerDetails)
		return
	}

	// create a new player from the body
	player := &models.Player{
		Name:        body.Name,
		FlagID:      body.FlagID,
		FargoID:     body.FargoID,
		FargoRating: body.FargoRating,
	}

	// add new player to the database
	if err := player.Create(server.db); err != nil {
		server.handleError(w, r, http.StatusInternalServerError, err)
		return
	}

	// return results
	server.handleSuccess(w, r, player)
}

// Handler for /players/count.
func (server *Server) handlePlayersCount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handlePlayersCountGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Players count handler for GET method.
func (server *Server) handlePlayersCountGet(w http.ResponseWriter, r *http.Request) {
	// get count of players
	var count int64
	countResult := server.db.Model(&models.Player{}).Count(&count)
	if countResult.Error != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	// return results
	server.handleSuccess(w, r, &CountResp{
		Count: count,
	})
}

// Handler for /players/{playerID}.
func (server *Server) handlePlayerByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handlePlayerByIDGet(w, r)
		case "PATCH":
			server.handlePlayerByIDPatch(w, r)
		case "DELETE":
			server.handlePlayerByIDDelete(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// PlayerById handler for GET method.
func (server *Server) handlePlayerByIDGet(w http.ResponseWriter, r *http.Request) {
	// get param for player id from url
	params := mux.Vars(r)
	playerID, ok := params["playerID"]
	if !ok || len(playerID) < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPlayerID)
		return
	}

	// convert player id to int
	id, err := strconv.Atoi(playerID)
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPlayerID)
		return
	}

	// get player by id
	var player models.Player
	if err := player.LoadByID(server.db, id); err != nil {
		if err == models.ErrPlayerIDInvalid {
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			return
		}
		if err == models.ErrPlayerNotFound {
			server.handleError(w, r, http.StatusNotFound, ErrPlayerNotFound)
			return
		}
	}

	// return results
	server.handleSuccess(w, r, player)
}

// PlayerByID handler for PATCH method. Updates an existing player.
func (server *Server) handlePlayerByIDPatch(w http.ResponseWriter, r *http.Request) {
	// get param for player id from url
	params := mux.Vars(r)
	playerID, ok := params["playerID"]
	if !ok || len(playerID) < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPlayerID)
		return
	}

	// convert player id to int
	id, err := strconv.Atoi(playerID)
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPlayerID)
		return
	}

	// decode the body
	var body PlayersPatchBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPlayerDetails)
		return
	}

	// get player by id
	var player models.Player
	if err := player.LoadByID(server.db, id); err != nil {
		if err == models.ErrPlayerIDInvalid {
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			return
		}
		if err == models.ErrPlayerNotFound {
			server.handleError(w, r, http.StatusNotFound, ErrPlayerNotFound)
			return
		}
	}

	// update player details from the body
	player.Name = body.Name
	player.FlagID = body.FlagID
	player.FargoID = body.FargoID
	player.FargoRating = body.FargoRating

	// update player in the database
	if err := player.Update(server.db); err != nil {
		server.handleError(w, r, http.StatusInternalServerError, err)
		return
	}

	// return results
	server.handleSuccess(w, r, player)
}

// PlayerByID handler for DELETE method. Deletes an existing player.
func (server *Server) handlePlayerByIDDelete(w http.ResponseWriter, r *http.Request) {
	// get param for player id from url
	params := mux.Vars(r)
	playerID, ok := params["playerID"]
	if !ok || len(playerID) < 1 {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPlayerID)
		return
	}

	// convert player id to int
	id, err := strconv.Atoi(playerID)
	if err != nil {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPlayerID)
		return
	}

	// get player by id
	var player models.Player
	if err := player.LoadByID(server.db, id); err != nil {
		if err == models.ErrPlayerIDInvalid {
			server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
			return
		}
		if err == models.ErrPlayerNotFound {
			server.handleError(w, r, http.StatusNotFound, ErrPlayerNotFound)
			return
		}
	}

	// delete player from database
	if err := player.Delete(server.db); err != nil {
		server.handleError(w, r, http.StatusInternalServerError, err)
		return
	}

	// return 204
	server.handle204Success(w, r)
}
