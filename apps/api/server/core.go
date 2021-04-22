package main

import (
	"fmt"
	"os"

	"github.com/codephobia/pool-overlay/apps/api/pkg/api"
	"github.com/codephobia/pool-overlay/apps/api/pkg/overlay"
	"github.com/codephobia/pool-overlay/apps/api/pkg/state"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Core is the main application.
type Core struct {
	db      *gorm.DB
	server  *api.Server
	overlay *overlay.Overlay
	state   *state.State
}

// NewCore returns a new Core.
func NewCore() (*Core, error) {
	// Connect to database.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Initialize Overlay.
	overlay := overlay.NewOverlay()

	// Initialize game state.
	state := state.NewState()

	// Initialize API Server.
	apiConfig := &api.Config{
		Host:      "0.0.0.0",
		Port:      "1268",
		PublicDir: "public",
		Version: &api.Version{
			Current:  "1",
			Previous: "1",
		},
	}
	server := api.NewServer(apiConfig, db, overlay, state)

	return &Core{
		db:      db,
		server:  server,
		overlay: overlay,
		state:   state,
	}, nil
}

// Init initializes components in the core.
func (c *Core) Init() {
	// Initialize Overlay.
	go c.overlay.Init()

	// Initialize API Server.
	c.server.Init()
}
