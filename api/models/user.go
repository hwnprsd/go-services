package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint      `json:"id,omitempty"	gorm:"primaryKey;autoIncrement"`
	Email     *string   `json:"email,omitempty"	gorm:"unique"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func MigrateUsers(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}

	return nil
}
