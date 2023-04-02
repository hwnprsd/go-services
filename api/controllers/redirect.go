package controllers

import (
	"errors"
	"log"
	"strings"

	"flaq.club/api/models"
	"github.com/gofiber/fiber/v2"
	fiberw "github.com/hwnprsd/go-easy-docs/fiber-wrapper"
)

func (c *Controller) SetupRedirectRoutes() {
	group := fiberw.NewGroup(c.FiberApp, "/r")
	fiberw.RawGet(group, "/", c.Redirect).WithQuery("url")
}

func (c *Controller) Redirect(ctx *fiber.Ctx) error {
	url := ctx.Query("url")
	utm := ctx.Query("utm")
	if !strings.Contains(url, "flaq.club") {
		log.Println("Invalid Redirection to -", url)
		err := "Invalid redirection request"
		return fiberw.NewRequestError(400, err, errors.New(err))
	}
	redirection := models.Redirect{
		Url: url,
		Utm: utm,
	}
	c.DB.Create(&redirection)
	log.Println("Redirecting to -", url)
	return ctx.Redirect(url)
}
