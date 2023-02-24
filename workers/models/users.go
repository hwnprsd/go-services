package models

import (
	"time"

	"gorm.io/gorm"
)

type MailingUser struct {
	ID           uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	EmailAddress string         `gorm:"primarykey" json:"email_address"`
	ListName     string         `gorm:"primarykey" json:"list_name"`
	Level        uint           `json:"level"`
}
