package models

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrRaceToInvalid  = errors.New("invalid race to amount")
	ErrRaceToNotFound = errors.New("handicap for race to not found")
)

// FargoHotHandicaps are an array of FargoHotHandicap.
type FargoHotHandicaps []FargoHotHandicap

// FargoHotHandicap is a handicap for the specified race to.
type FargoHotHandicap struct {
	ID              uint  `json:"id" gorm:"primarykey"`
	RaceTo          int   `json:"race_to" gorm:"index"`
	DifferenceStart uint  `json:"difference_start"`
	DifferenceEnd   *uint `json:"difference_end"`
	WinsHigher      uint  `json:"wins_higher"`
	WinsLower       uint  `json:"wins_lower"`
}

// LoadByRaceTo loads handicaps for the supplied race to.
func (f *FargoHotHandicaps) LoadByRaceTo(database *gorm.DB, raceTo int) error {
	if raceTo < minHandicapRaceTo || raceTo > maxHandicapRaceTo {
		return ErrRaceToInvalid
	}

	result := database.
		Select("id", "race_to", "difference_start", "difference_end", "wins_higher", "wins_lower").
		Where("race_to = ?", raceTo).
		Order("difference_start asc").
		Find(f)

	if result.Error != nil {
		return ErrRaceToInvalid
	}

	if result.RowsAffected == 0 {
		return ErrRaceToNotFound
	}

	return nil
}
