package main

import (
	"log"

	"github.com/joho/godotenv"

	_ "github.com/codephobia/pool-overlay/libs/go/apidocs"
)

func main() {
	// Load .env file. This is only used for local running outside of docker,
	// which is why we ignore the error. In docker, we add the env vars via
	// docker-compose which points to the same .env file.
	if err := godotenv.Load(); err != nil {
		log.Printf("[INFO] skipping loading .env file")
	}

	// Create new core.
	core, err := NewCore()
	if err != nil {
		log.Fatalf("[FATAL] core: %s", err)
	}

	// Initialize core.
	core.Init()

	// Start API server.
	if err := core.server.Run(); err != nil {
		log.Fatalf("[FATAL] server: %s", err)
	}
}
