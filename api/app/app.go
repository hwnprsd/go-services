package app

import (
	"flaq.club/api/messaging"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
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

func (a *App) HealthCheck(c *fiber.Ctx) error {
	c.SendString("Health Check from Go")
	return nil
}

func (a *App) SetupRoutes() {
	a.FiberApp.Get("/health", a.HealthCheck)

	a.FiberApp.Get("/test", func(c *fiber.Ctx) error {
		c.SendString("Test Func 2")
		return nil
	})

	a.FiberApp.Get("/mailer", func(c *fiber.Ctx) error {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Ashwin is a genius lmao"),
		}
		if err := a.MQ.MailerQueue.Channel.Publish(
			"",       // exchange
			"Mailer", // queue name
			false,    // mandatory
			false,    // immediate
			message,  // message to publish
		); err != nil {
			return err
		}
		c.SendString("Message Published to AMQP")
		return nil
	})
}
