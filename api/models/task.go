package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Category  string         `json:"category"`
	Status    string         `json:"status"`
	UserID    uint           `json:"user_id"`
	User      *User          `json:"user"`
}
