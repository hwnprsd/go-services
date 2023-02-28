package models

import (
	"time"

	"gorm.io/gorm"
)

type MailingList struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	ListName  string         `gorm:"uniqueindex" json:"event_name"`
	Users     []*MailingUser `gorm:"many2many:mailer_user_list" json:"users"`
}
