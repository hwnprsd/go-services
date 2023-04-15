package models

import (
	"time"
)

type SummaryNewsletter struct {
	ID          uint      `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Articles    []Article `json:"articles" gorm:"many2many:summary_newsletter_article"`
	Summary     string    `json:"summary"`
	PublishDate time.Time `json:"publish_date"`
	TagID       uint      `json:"tag_id"`
	Tag         Tag       `json:"tag"`
}
