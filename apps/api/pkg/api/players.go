package api

import (
	"math"
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/apps/api/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const (
	playersPerPage = 10
)

// Handler for /players.
func (server *Server) handlePlayers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			server.handlePlayersGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Players handler for GET method.
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
		Select("id", "name").
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

// Handler for /players/count.
func (server *Server) handlePlayersCount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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
		case "GET":
			server.handlePlayerByIDGet(w, r)
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
