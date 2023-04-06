package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt         time.Time      `json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Email             *string        `json:"email,omitempty" gorm:"uniqueIndex"`
	UniqueId          string         `json:"unique_id" gorm:"uniqueIndex"`
	Level             uint           `json:"level"`
	Name              string         `json:"name"`
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	NickName          string         `json:"nick_name"`
	Description       string         `json:"description"`
	UserID            string         `json:"user_id"`
	AvatarURL         string         `json:"avatar_url"`
	Location          string         `json:"location"`
	AccessTokenSecret string         `json:"access_token_secret"`
	RefreshToken      string         `json:"refresh_token"`
	ExpiresAt         time.Time      `json:"expires_at"`
}
