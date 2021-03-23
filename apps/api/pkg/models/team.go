package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrTeamIDInvalid = errors.New("invalid team id")
	ErrTeamNotFound  = errors.New("team not found")
)

// Team is a team of pool players.
type Team struct {
	ID      uint          `json:"id,omitempty" gorm:"primarykey"`
	Name    string        `json:"name,omitempty" gorm:"size:100"`
	Players []*TeamPlayer `json:"players,omitempty" gorm:"foreignKey:team_id"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// LoadByID loads a team by ID.
func (t *Team) LoadByID(database *gorm.DB, id int) error {
	result := database.
		Select("id", "name").
		Where("id = ?", id).
		Preload("Players", func(db *gorm.DB) *gorm.DB {
			return db.
				Select("id", "team_id", "player_id", "captain").
				Preload("Player", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "name")
				})
		}).
		Find(t)

	if result.Error != nil {
		return ErrTeamIDInvalid
	}

	if result.RowsAffected != 1 {
		return ErrTeamNotFound
	}

	return nil
}
