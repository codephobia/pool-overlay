package seeds

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"gorm.io/gorm"

	"github.com/codephobia/pool-overlay/libs/go/models"
)

// FlagSeed is used to load all flag seeds from the json file.
type FlagSeed struct {
	Flags []models.Flag
}

// NewFlagSeed return a new FlagSeed.
func NewFlagSeed(jsonPath ...string) *FlagSeed {
	flagSeed := &FlagSeed{}

	if err := flagSeed.load(jsonPath...); err != nil {
		panic(err)
	}

	return flagSeed
}

// Load loads the json file.
func (f *FlagSeed) load(jsonPath ...string) error {
	seedFile, err := os.Open(path.Join(jsonPath...))
	if err != nil {
		return err
	}
	defer seedFile.Close()

	return json.NewDecoder(seedFile).Decode(&f.Flags)
}

// Seeds returns all flags loaded from the json file as seeds.
func (f *FlagSeed) Seeds() []Seed {
	seeds := make([]Seed, 0)

	for _, flag := range f.Flags {
		newSeed := Seed{
			Name: "CreateFlag" + flag.Country,
			Run: (func(flag models.Flag) func(db *gorm.DB) error {
				return func(db *gorm.DB) error {
					return CreateFlag(db, flag.ID, flag.Country, flag.ImagePath)
				}
			})(flag),
		}

		seeds = append(seeds, newSeed)
	}

	return seeds
}

// CreateFlag adds a flag to the database.
func CreateFlag(db *gorm.DB, id uint, country string, imagePath string) error {
	var flags []models.Flag
	db.Where("id = ?", id).Find(&flags)

	if len(flags) == 1 {
		log.Printf("flag already exists for country: %s", country)
		return nil
	}

	return db.Create(
		&models.Flag{
			ID:        id,
			Country:   country,
			ImagePath: imagePath,
		},
	).Error
}
