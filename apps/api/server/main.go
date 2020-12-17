package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	core, err := NewCore()
	if err != nil {
		log.Fatalf("core: %s", err)
	}
	core.Run()

	<-core.done
}
