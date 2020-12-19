package models

import (
	"time"

	"gorm.io/gorm"
)

// TeamPlayer is a player on a team.
type TeamPlayer struct {
	ID       uint `json:"id,omitempty" gorm:"primarykey"`
	TeamID   uint `json:"team_id,omitempty"`
	PlayerID uint `json:"player_id,omitempty"`
	Captain  bool `json:"captain,omitempty"`

	Team   *Team   `json:"team,omitempty"`
	Player *Player `json:"player,omitempty"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
