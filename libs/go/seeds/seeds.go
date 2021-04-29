package seeds

import (
	"fmt"

	"gorm.io/gorm"
)

// Seed is a database entry to be added.
type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

// Seeder is used to load a json seed file and return the seed slice.
type Seeder interface {
	Seeds() []Seed
}

// Run takes a variadic number of loaders and runs all of their seeds.
func Run(db *gorm.DB, seeders ...Seeder) error {
	seeds := make([]Seed, 0)

	for _, seeder := range seeders {
		s := seeder.Seeds()

		seeds = append(seeds, s...)
	}

	for _, seed := range seeds {
		if err := seed.Run(db); err != nil {
			return fmt.Errorf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	return nil
}
