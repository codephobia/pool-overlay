package seeds

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"gorm.io/gorm"

	"github.com/codephobia/pool-overlay/libs/go/models"
)

// PlayerSeed is used to load all player seeds from the json file.
type PlayerSeed struct {
	Players []models.Player
}

// NewPlayerSeed return a new PlayerSeed.
func NewPlayerSeed(jsonPath ...string) *PlayerSeed {
	playerSeed := &PlayerSeed{}

	if err := playerSeed.load(jsonPath...); err != nil {
		panic(err)
	}

	return playerSeed
}

// Loads the json file.
func (s *PlayerSeed) load(jsonPath ...string) error {
	seedFile, err := os.Open(path.Join(jsonPath...))
	if err != nil {
		return err
	}
	defer seedFile.Close()

	return json.NewDecoder(seedFile).Decode(&s.Players)
}

// Seeds returns all flags loaded from the json file as seeds.
func (s *PlayerSeed) Seeds() []Seed {
	seeds := make([]Seed, 0)

	for _, player := range s.Players {
		newSeed := Seed{
			Name: "CreatePlayer" + player.Name,
			Run: (func(player models.Player) func(db *gorm.DB) error {
				return func(db *gorm.DB) error {
					return CreatePlayer(db, player.ID, player.Name, player.FlagID, player.FargoID, player.FargoRating)
				}
			})(player),
		}

		seeds = append(seeds, newSeed)
	}

	return seeds
}

// CreatePlayer adds a player to the database.
func CreatePlayer(db *gorm.DB, id uint, name string, flagID uint, fargoID uint, fargoRating uint) error {
	var players []models.Player
	db.Where("id = ?", id).Find(&players)

	if len(players) == 1 {
		log.Printf("player already exists: %s", name)
		return nil
	}

	return db.Create(
		&models.Player{
			ID:          id,
			Name:        name,
			FlagID:      flagID,
			FargoID:     fargoID,
			FargoRating: fargoRating,
		},
	).Error
}
