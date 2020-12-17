package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codephobia/pool-overlay/apps/api/pkg/api"
	"github.com/codephobia/pool-overlay/apps/api/pkg/overlay"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Core is the main application.
type Core struct {
	db      *gorm.DB
	api     *api.API
	overlay *overlay.Overlay

	done chan struct{}
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

	// Initialize API Server.
	apiConfig := &api.Config{
		Scheme:    "https",
		Host:      "0.0.0.0",
		Port:      "1268",
		PublicDir: "public",
	}
	api := api.NewAPI(apiConfig, overlay)

	return &Core{
		db:      db,
		api:     api,
		overlay: overlay,

		done: make(chan struct{}),
	}, nil
}

// Run starts the main application.
func (c *Core) Run() error {
	log.Printf("running")

	// Initialize Overlay.
	go c.overlay.Init()

	// Initialize API Server.
	return c.api.Init()
}
