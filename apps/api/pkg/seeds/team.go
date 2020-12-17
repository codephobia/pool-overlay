package seeds

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"gorm.io/gorm"

	"github.com/codephobia/pool-overlay/apps/api/pkg/models"
)

// TeamSeed is used to load all player seeds from the json file.
type TeamSeed struct {
	Teams []models.Team
}

// NewTeamSeed return a new TeamSeed.
func NewTeamSeed(jsonPath ...string) *TeamSeed {
	teamSeed := &TeamSeed{}

	if err := teamSeed.load(jsonPath...); err != nil {
		panic(err)
	}

	return teamSeed
}

// Loads the json file.
func (s *TeamSeed) load(jsonPath ...string) error {
	seedFile, err := os.Open(path.Join(jsonPath...))
	if err != nil {
		return err
	}
	defer seedFile.Close()

	return json.NewDecoder(seedFile).Decode(&s.Teams)
}

// Seeds returns all teams loaded from the json file as seeds.
func (s *TeamSeed) Seeds() []Seed {
	seeds := make([]Seed, 0)

	for _, team := range s.Teams {
		runFunc := (func(team models.Team) func(db *gorm.DB) error {
			return func(db *gorm.DB) error {
				return CreateTeam(db, team.ID, team.Name)
			}
		})(team)

		newSeed := Seed{
			Name: "CreateTeam" + team.Name,
			Run:  runFunc,
		}

		seeds = append(seeds, newSeed)
	}

	return seeds
}

// CreateTeam adds a team to the database.
func CreateTeam(db *gorm.DB, id uint, name string) error {
	var teams []models.Team
	db.Where("id = ?", id).Find(&teams)

	if len(teams) == 1 {
		log.Printf("team already exists: %s", name)
		return nil
	}

	return db.Create(
		&models.Team{
			ID:   id,
			Name: name,
		},
	).Error
}
