package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codephobia/pool-overlay/apps/api/pkg/models"
	"github.com/codephobia/pool-overlay/apps/api/pkg/seeds"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load env vars.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

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
		panic(err)
	}

	// Make sure tables, indexes, etc are all created.
	db.AutoMigrate(
		&models.Flag{},
		&models.Player{},
		&models.Team{},
		&models.TeamPlayer{},
	)

	// Load seeds from json files.
	if err := seeds.Run(
		db,
		seeds.NewFlagSeed("data", "flags.json"),
		seeds.NewPlayerSeed("data", "players.json"),
		seeds.NewTeamSeed("data", "teams.json"),
		seeds.NewTeamPlayerSeed("data", "team_players.json"),
	); err != nil {
		panic(err)
	}

	log.Println("done")
}
