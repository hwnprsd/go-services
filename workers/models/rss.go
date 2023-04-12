package models

import (
	"time"

	"gorm.io/gorm"
)

type RssFeed struct {
	ID            uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	LastRefreshed time.Time      `json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Url           string         `json:"url"`
	Tags          []Tag          `json:"tags" gorm:"many2many:rss_tags;"`
	ShouldSkip    bool           `json:"should_skip"`
}
