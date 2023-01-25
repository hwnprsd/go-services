package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type EmailTemplate struct {
	ID               uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt        time.Time      `json:"created_at,omitempty"`
	UpdatedAt        time.Time      `json:"updated_at,omitempty"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	TemplateString   string         `json:"template_string,omitempty"`
	ValueArrangement pq.StringArray `json:"value_arrangement" gorm:"type:text[]"`
}
