package controllers

import (
	"log"

	"flaq.club/api/app"
	"flaq.club/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Controller struct {
	*app.App
}

func New() *fiber.App {
	fiberApp := fiber.New()
	fiberApp.Use(logger.New())
	return fiberApp
}

const API_UPDATE_VERSION = 1

func (c *Controller) SetupRoutes() {
	log.Println("Setting up routes")
	c.FiberApp.Get("/health", func(ctx *fiber.Ctx) error {
		ctx.JSON(fiber.Map{
			"success":     true,
			"api_version": API_UPDATE_VERSION,
		})
		return nil
	})

	c.FiberApp.Post("/users/create", func(ct *fiber.Ctx) error {
		body := new(CreateUserBody)
		return utils.PostRequestHandler(ct, c.CreateUser(), utils.RequestBody{Data: body})
	})

	c.FiberApp.Get("/users", func(ctx *fiber.Ctx) error {
		return utils.GetRequestHandler(ctx, c.GetUsers(), utils.RequestBody{})
	})
}
