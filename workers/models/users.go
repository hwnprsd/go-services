package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MailingUser struct {
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	EmailAddress  string         `gorm:"primarykey" json:"email_address"`
	Level         uint           `json:"level"`
	Name          string         `json:"name"`
	WalletAddress string         `json:"wallet_address"`
	Lists         []*MailingList `gorm:"many2many:mailer_user_list" json:"lists"`
	MetaData      datatypes.JSON `json:"meta_data"`
}
