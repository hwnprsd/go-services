package models

import (
	"time"

	"gorm.io/gorm"
)

// TODO Make this more dynamic in terms of handling dynamic templates
type MailingCampaign struct {
	ID              uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	ListName        string         `json:"list_name"`
	IsRecurring     bool           `json:"is_recurring"`
	EmailSourceType uint           `json:"email_source_type"` // 1 = Dynamic, 2 = Template String
	TemplateString  string         `json:"template_string"`
	TemplateURL     string         `json:"template_url"`
}
