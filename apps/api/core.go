package main

import (
	"fmt"
	"os"

	"github.com/codephobia/pool-overlay/libs/go/api"
	overlayPkg "github.com/codephobia/pool-overlay/libs/go/overlay"
	"github.com/codephobia/pool-overlay/libs/go/state"
	"github.com/codephobia/pool-overlay/libs/go/telestrator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Core is the main application.
type Core struct {
	db          *gorm.DB
	server      *api.Server
	overlay     *overlayPkg.Overlay
	telestrator *telestrator.Telestrator
	tables      map[int]*state.State
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

	// Initialize Scoreboard Overlay.
	overlay := overlayPkg.NewOverlay()

	// Initialize Telestrator Overlay.
	telestrator := telestrator.NewTelestrator()

	// Initialize game state.
	tables := map[int]*state.State{}
	tables[1] = state.NewState(db, 1)
	tables[2] = state.NewState(db, 2)

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
	server := api.NewServer(apiConfig, db, overlay, telestrator, tables)

	return &Core{
		db:          db,
		server:      server,
		overlay:     overlay,
		telestrator: telestrator,
		tables:      tables,
	}, nil
}

// Init initializes components in the core.
func (c *Core) Init() {
	// Initialize Scoreboard Overlay.
	go c.overlay.Init()

	// Initialize Telestrator Overlay.
	go c.telestrator.Overlay.Init()

	// Initialize API Server.
	c.server.Init()
}
