package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrFlagIDInvalid = errors.New("invalid flag id")
	ErrFlagNotFound  = errors.New("flag not found")
)

// Flag is a country flag for a pool player.
type Flag struct {
	ID        uint   `json:"id,omitempty" gorm:"primarykey"`
	Country   string `json:"country,omitempty" gorm:"size:100"`
	ImagePath string `json:"image_path,omitempty" gorm:"size:100"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// LoadByID loads a flag by ID.
func (f *Flag) LoadByID(database *gorm.DB, id int) error {
	result := database.
		Select("id", "country", "image_path").
		Where("id = ?", id).
		Find(f)

	if result.Error != nil {
		return ErrFlagIDInvalid
	}

	if result.RowsAffected != 1 {
		return ErrFlagNotFound
	}

	return nil
}
