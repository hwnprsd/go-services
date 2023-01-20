package app

import (
	"flaq.club/api/messaging"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	DB       *gorm.DB
	MQ       *messaging.Messaging
	FiberApp *fiber.App
}

func New(m *messaging.Messaging, f *fiber.App, d *gorm.DB) *App {

	defaultApp := &App{
		MQ:       m,
		FiberApp: f,
		DB:       d,
	}
	return defaultApp
}
