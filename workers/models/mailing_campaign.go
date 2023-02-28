package models

import (
	"time"

	"gorm.io/gorm"
)

// TODO Make this more dynamic in terms of handling dynamic templates
type MailingCampaign struct {
	ID                      uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt               time.Time      `json:"created_at,omitempty"`
	UpdatedAt               time.Time      `json:"updated_at,omitempty"`
	DeletedAt               gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	MailingListID           uint           `json:"mailing_list_id"`
	MailingList             MailingList    `json:"mailing_list"`
	EmailTemplateID         uint           `json:"email_template_id"`
	EmailTemplate           EmailTemplate  `json:"email_template"`
	IsRecurring             bool           `json:"is_recurring"`
	MetaKey                 string         `json:"meta_key"`
	MetaValue               uint           `json:"meta_value"`
	ShouldIncrementMeta     bool           `json:"should_increment_meta"`
	ShouldInitMeta          bool           `json:"should_init_meta"`
	ShouldInjectUserDetails bool           `json:"should_inject_user_details"`
}
