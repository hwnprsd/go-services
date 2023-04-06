package controllers

import (
	"log"

	"flaq.club/api/app"
	"github.com/MicahParks/keyfunc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	fiberw "github.com/hwnprsd/go-easy-docs/fiber-wrapper"
)

type Controller struct {
	*app.App
	jwk *keyfunc.JWKS
}

func New() *fiber.App {
	engine := html.New("./views", ".html")
	fiberApp := fiber.New(fiber.Config{
		Views: engine,
	})
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

	c.SetupAuthRoutes()
	c.SetupNFTRoutes()
	c.SetupQuizRoutes()
	c.SetupUserRoutes()
	c.SetupTaskRoutes()
	c.SetupRedirectRoutes()

	c.FiberApp.Get("/api/docs", func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		fiberw.WriteApiDocumentation("Flaq Academy API Service", "Flaq helps provide web3 education with EASE", ctx.Response().BodyWriter())
		return nil
	})
}
