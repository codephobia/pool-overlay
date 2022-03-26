package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrPlayerIDInvalid    = errors.New("invalid player id")
	ErrPlayerNotFound     = errors.New("player not found")
	ErrPlayerHasIDAlready = errors.New("player already has an id")
)

// Player is a pool player.
type Player struct {
	ID          uint   `json:"id,omitempty" gorm:"primarykey"`
	Name        string `json:"name,omitempty" gorm:"size:100"`
	FlagID      uint   `json:"flag_id,omitempty"`
	Flag        *Flag  `json:"flag,omitempty" gorm:"foreignKey:flag_id"`
	FargoID     uint   `json:"fargo_id,omitempty"`
	FargoRating uint   `json:"fargo_rating,omitempty"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// LoadByID loads a player by ID.
func (p *Player) LoadByID(database *gorm.DB, id int) error {
	result := database.
		Select("id", "name", "flag_id", "fargo_id", "fargo_rating").
		Where("id = ?", id).
		Preload("Flag", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "country", "image_path")
		}).
		Find(p)

	// could use:
	// errors.Is(err, gorm.ErrRecordNotFound)

	if result.Error != nil {
		return ErrPlayerIDInvalid
	}

	if result.RowsAffected != 1 {
		return ErrPlayerNotFound
	}

	return nil
}

// Create adds the player to the database.
func (p *Player) Create(database *gorm.DB) error {
	if p.ID != 0 {
		return ErrPlayerHasIDAlready
	}

	return database.Create(p).Error
}

// Update updates the current player in the database.
func (p *Player) Update(database *gorm.DB) error {
	if p.ID == 0 {
		return ErrPlayerIDInvalid
	}

	return database.Model(p).Updates(map[string]interface{}{
		"name":         p.Name,
		"flag_id":      p.FlagID,
		"fargo_id":     p.FargoID,
		"fargo_rating": p.FargoRating,
	}).Error
}

// Delete removes the current player from the database.
func (p *Player) Delete(database *gorm.DB) error {
	if p.ID == 0 {
		return ErrPlayerIDInvalid
	}

	return database.Delete(p).Error
}
