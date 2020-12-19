package models

import (
	"time"

	"gorm.io/gorm"
)

// Flag is a country flag for a pool player.
type Flag struct {
	ID        uint   `json:"id,omitempty" gorm:"primarykey"`
	Country   string `json:"country,omitempty" gorm:"size:100"`
	ImagePath string `json:"image_path,omitempty" gorm:"size:100"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
