package models

import "gorm.io/gorm"

// Flag is a country flag for a pool player.
type Flag struct {
	ID        uint   `json:"id"`
	Country   string `json:"country" gorm:"size:100"`
	ImagePath string `json:"image_path" gorm:"size:100"`

	gorm.Model
}
