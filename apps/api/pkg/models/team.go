package models

import "gorm.io/gorm"

// Team is a team of pool players.
type Team struct {
	ID      uint         `json:"id"`
	Name    string       `json:"name" gorm:"size:100"`
	Players []TeamPlayer `gorm:"foreignKey:team_id"`

	gorm.Model
}
