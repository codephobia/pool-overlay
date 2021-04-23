package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrPlayerIDInvalid = errors.New("invalid player id")
	ErrPlayerNotFound  = errors.New("player not found")
)

// Player is a pool player.
type Player struct {
	ID     uint   `json:"id,omitempty" gorm:"primarykey"`
	Name   string `json:"name,omitempty" gorm:"size:100"`
	FlagID uint   `json:"flag_id,omitempty"`
	Flag   *Flag  `json:"flag,omitempty" gorm:"foreignKey:flag_id"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// LoadByID loads a player by ID.
func (p *Player) LoadByID(database *gorm.DB, id int) error {
	result := database.
		Select("id", "name", "flag_id").
		Where("id = ?", id).
		Preload("Flag", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "country", "image_path")
		}).
		Find(p)

	if result.Error != nil {
		return ErrPlayerIDInvalid
	}

	if result.RowsAffected != 1 {
		return ErrPlayerNotFound
	}

	return nil
}
