package models

import (
	"time"

	"gorm.io/gorm"
)

type Redirect struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Url       string         `json:"url"`
	Utm       string         `json:"utm"`
}
