package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID          uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	PublishDate time.Time      `json:"publish_date"`
	Title       string         `json:"title"`
	Summary     string         `json:"summary"`
	Authors     string         `json:"authors"`
	Url         string         `json:"link"`
	Tag         []Tag          `json:"tags" gorm:"many2many:articles_tags;"`
	GUID        string         `json:"guid" gorm:"uniqueIndex:article_guid;null"`
}
