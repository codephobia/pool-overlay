package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create new core.
	core, err := NewCore()
	if err != nil {
		log.Fatalf("core: %s", err)
	}

	// Initialize core.
	core.Init()

	// Start API server.
	if err := core.server.Run(); err != nil {
		log.Fatalf("server: %s", err)
	}
}
