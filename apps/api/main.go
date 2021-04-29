package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file. This is only used for local running outside of docker,
	// which is why we ignore the error. In docker, we add the env vars via
	// docker-compose which points to the same .env file.
	godotenv.Load()

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
