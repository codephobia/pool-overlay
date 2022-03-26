package api

import (
	"math"
	"net/http"
	"strconv"

	"github.com/codephobia/pool-overlay/libs/go/models"
	"gorm.io/gorm"
)

const (
	gamesPerPage = 100
)

// Handler for /games.
func (server *Server) handleGames() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleGamesGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Games handler for GET method. Returns a page of completed games.
func (server *Server) handleGamesGet(w http.ResponseWriter, r *http.Request) {
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

	// TODO: add game filtering

	// get count of completed games to test page ceiling
	var count int64
	countResult := server.db.Model(&models.Game{}).
		Where("completed = ?", true).
		Count(&count)
	if countResult.Error != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	// check if page is beyond maximum
	totalPages := int(math.Ceil(float64(count) / gamesPerPage))
	if pageNum > totalPages {
		server.handleError(w, r, http.StatusUnprocessableEntity, ErrInvalidPageNumber)
		return
	}

	// get games for page
	var games []*models.Game

	offset := pageNum*gamesPerPage - gamesPerPage
	result := server.db.
		Select(
			"id",
			"type",
			"vs_mode",
			"race_to",
			"score_one",
			"score_two",
			"completed",
			"player_one_id",
			"player_two_id",
			"use_fargo_hot_handicap",
			"updated_at",
		).
		Where("completed = ?", true).
		Order("updated_at DESC").
		Limit(gamesPerPage).
		Offset(offset).
		Preload("PlayerOne", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("PlayerTwo", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Find(&games)
	if result.Error != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	// return results
	server.handleSuccess(w, r, games)
}

// Handler for /games/count.
func (server *Server) handleGamesCount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			server.HandleOptions(w, r)
		case "GET":
			server.handleGamesCountGet(w, r)
		default:
			server.handleError(w, r, http.StatusMethodNotAllowed, ErrEndpointMethodNotAllowed)
		}
	})
}

// Games count handler for GET method.
func (server *Server) handleGamesCountGet(w http.ResponseWriter, r *http.Request) {
	// TODO: add game filtering

	// get count of games
	var count int64
	countResult := server.db.Model(&models.Game{}).
		Where("completed = ?", true).
		Count(&count)
	if countResult.Error != nil {
		server.handleError(w, r, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	// return results
	server.handleSuccess(w, r, &CountResp{
		Count: count,
	})
}
