package models

import (
	"time"

	"gorm.io/gorm"
)

// Player is a pool player.
type Player struct {
	ID     uint   `json:"id,omitempty" gorm:"primarykey"`
	Name   string `json:"name,omitempty" gorm:"size:100"`
	FlagID uint   `json:"flag_id,omitempty"`
	Flag   *Flag  `json:"flag,omitempty" gorm:"foreignKey:id"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
