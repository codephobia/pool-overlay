package models

import "gorm.io/gorm"

// TeamPlayer is a player on a team.
type TeamPlayer struct {
	ID       uint `json:"id"`
	TeamID   uint `json:"team_id"`
	PlayerID uint `json:"player_id"`
	Captain  bool `json:"captain"`

	Team   Team   `json:"team"`
	Player Player `json:"player"`

	gorm.Model
}
