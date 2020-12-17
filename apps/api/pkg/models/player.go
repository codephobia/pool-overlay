package models

import "gorm.io/gorm"

// Player is a pool player.
type Player struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" gorm:"size:100"`
	FlagID uint   `json:"flag_id"`
	Flag   Flag   `json:"flag" gorm:"foreignKey:id"`

	gorm.Model
}
