package models

import (
	"time"

	"gorm.io/gorm"
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
