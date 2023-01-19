package database

import (
	"fmt"
	"log"
	"os"

	"flaq.club/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	PG_HOST := os.Getenv("PG_HOST")
	PG_USERNAME := os.Getenv("PG_USERNAME")
	PG_PASSWORD := os.Getenv("PG_PASSWORD")
	PG_PORT := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable TimeZone=Asia/Kolkata", PG_HOST, PG_USERNAME, PG_PASSWORD, PG_PORT)
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	dbInstance.AutoMigrate(&models.User{})
	return dbInstance
}
