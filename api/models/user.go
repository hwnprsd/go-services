package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Email     *string        `json:"email,omitempty" gorm:"uniqueIndex"`
	Level     uint           `json:"level"`
}

func MigrateUsers(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Users Migrated")
	return nil
}
