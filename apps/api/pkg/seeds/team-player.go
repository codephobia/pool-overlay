package seeds

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"gorm.io/gorm"

	"github.com/codephobia/pool-overlay/apps/api/pkg/models"
)

// TeamPlayerSeed is used to load all team player seeds from the json file.
type TeamPlayerSeed struct {
	TeamPlayers []models.TeamPlayer
}

// NewTeamPlayerSeed return a new TeamPlayerSeed.
func NewTeamPlayerSeed(jsonPath ...string) *TeamPlayerSeed {
	teamPlayerSeed := &TeamPlayerSeed{}

	if err := teamPlayerSeed.load(jsonPath...); err != nil {
		panic(err)
	}

	return teamPlayerSeed
}

// Loads the json file.
func (s *TeamPlayerSeed) load(jsonPath ...string) error {
	seedFile, err := os.Open(path.Join(jsonPath...))
	if err != nil {
		return err
	}
	defer seedFile.Close()

	return json.NewDecoder(seedFile).Decode(&s.TeamPlayers)
}

// Seeds returns all teams loaded from the json file as seeds.
func (s *TeamPlayerSeed) Seeds() []Seed {
	seeds := make([]Seed, 0)

	for _, teamPlayer := range s.TeamPlayers {
		runFunc := (func(teamPlayer models.TeamPlayer) func(db *gorm.DB) error {
			return func(db *gorm.DB) error {
				return CreateTeamPlayer(db, teamPlayer.ID, teamPlayer.TeamID, teamPlayer.PlayerID)
			}
		})(teamPlayer)

		newSeed := Seed{
			Name: fmt.Sprintf("CreateTeamPlayer%d", teamPlayer.ID),
			Run:  runFunc,
		}

		seeds = append(seeds, newSeed)
	}

	return seeds
}

// CreateTeamPlayer adds a team player to the database.
func CreateTeamPlayer(db *gorm.DB, id uint, teamID uint, playerID uint) error {
	var teamPlayers []models.TeamPlayer
	db.Where("id = ?", id).Find(&teamPlayers)

	if len(teamPlayers) == 1 {
		log.Printf("team player already exists: %d", id)
		return nil
	}

	return db.Create(
		&models.TeamPlayer{
			ID:       id,
			TeamID:   teamID,
			PlayerID: playerID,
		},
	).Error
}
