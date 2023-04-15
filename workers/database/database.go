package database

import (
	"fmt"
	"log"
	"os"

	"flaq.club/workers/models"
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
	err = dbInstance.AutoMigrate(&models.EmailTemplate{}, &models.NftMint{}, &models.Web3Event{}, &models.MailingUser{}, &models.MailingList{}, &models.MailingCampaign{}, &models.Article{}, &models.RssFeed{}, &models.Tag{}, &models.SummaryNewsletter{})
	if err != nil {
		log.Fatal(err)
	}
	return dbInstance
}
