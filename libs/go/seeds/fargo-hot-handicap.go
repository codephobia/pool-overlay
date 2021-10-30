package seeds

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"strconv"

	"gorm.io/gorm"

	"github.com/codephobia/pool-overlay/libs/go/models"
)

// FargoHotHandicapSeed is used to load all handicap seeds from the json file.
type FargoHotHandicapSeed struct {
	Handicaps []models.FargoHotHandicap
}

// NewFargoHotHandicapSeed return a new FargoHotHandicapSeed.
func NewFargoHotHandicapSeed(jsonPath ...string) *FargoHotHandicapSeed {
	fargoHotHandicapSeed := &FargoHotHandicapSeed{}

	if err := fargoHotHandicapSeed.load(jsonPath...); err != nil {
		panic(err)
	}

	return fargoHotHandicapSeed
}

// Loads the json file.
func (s *FargoHotHandicapSeed) load(jsonPath ...string) error {
	seedFile, err := os.Open(path.Join(jsonPath...))
	if err != nil {
		return err
	}
	defer seedFile.Close()

	return json.NewDecoder(seedFile).Decode(&s.Handicaps)
}

// Seeds returns all handicaps loaded from the json file as seeds.
func (s *FargoHotHandicapSeed) Seeds() []Seed {
	seeds := make([]Seed, 0)

	for _, handicap := range s.Handicaps {
		id := strconv.Itoa(int(handicap.ID))
		newSeed := Seed{
			Name: "CreateFargoHotHandicap" + id,
			Run: (func(handicap models.FargoHotHandicap) func(db *gorm.DB) error {
				return func(db *gorm.DB) error {
					return CreateFargoHotHandicap(db, handicap.ID, handicap.RaceTo, handicap.DifferenceStart, handicap.DifferenceEnd, handicap.WinsHigher, handicap.WinsLower)
				}
			})(handicap),
		}

		seeds = append(seeds, newSeed)
	}

	return seeds
}

// CreateFargoHotHandicap adds a handicap to the database.
func CreateFargoHotHandicap(db *gorm.DB, id uint, raceTo int, differenceStart uint, differenceEnd *uint, winsHigher uint, winsLower uint) error {
	var handicaps []models.FargoHotHandicap
	db.Where("id = ?", id).Find(&handicaps)

	if len(handicaps) == 1 {
		log.Printf("fargo hot handicap id already exists: %d", int(id))
		return nil
	}

	return db.Create(
		&models.FargoHotHandicap{
			ID:              id,
			RaceTo:          raceTo,
			DifferenceStart: differenceStart,
			DifferenceEnd:   differenceEnd,
			WinsHigher:      winsHigher,
			WinsLower:       winsLower,
		},
	).Error
}
