package controllers

import (
	"log"

	"flaq.club/api/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Controller struct {
	*app.App
}

func New() *fiber.App {
	fiberApp := fiber.New()
	fiberApp.Use(logger.New())
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Update this
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	return fiberApp
}

const API_UPDATE_VERSION = 70

func (c *Controller) SetupRoutes() {
	log.Println("Setting up routes")
	c.FiberApp.Get("/health", func(ctx *fiber.Ctx) error {
		ctx.JSON(fiber.Map{
			"success":     true,
			"api_version": API_UPDATE_VERSION,
		})
		return nil
	})

	c.SetupNFTRoutes()
	c.SetupQuizRoutes()
	c.SetupUserRoutes()
	c.SetupTaskRoutes()
}
