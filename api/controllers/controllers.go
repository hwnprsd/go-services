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

const API_UPDATE_VERSION = 5

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

	c.FiberApp.Post("/nft/poap/mint", func(ctx *fiber.Ctx) error {
		body := new(MintNftBody)
		return utils.PostRequestHandler(ctx, c.MintPOAP(), utils.RequestBody{Data: body})
	})

	c.FiberApp.Post("/quiz/submit", func(ctx *fiber.Ctx) error {
		body := new(SubmitQuizParticipationBody)
		return utils.PostRequestHandler(ctx, c.SubmitQuizParticipation(), utils.RequestBody{Data: body})
	})

	c.FiberApp.Post("/quiz/request-email", func(ctx *fiber.Ctx) error {
		body := new(NFTClaimEmailBody)
		return utils.PostRequestHandler(ctx, c.RequestNFTClaimEmail(), utils.RequestBody{Data: body})
	})

	c.FiberApp.Get("/quiz/claim-info", func(ctx *fiber.Ctx) error {
		quizClaimId := ctx.Query("quizClaimId")
		log.Println(quizClaimId)
		query := new(GetSubmissionInfoQuery)
		query.QuizClaimID = quizClaimId
		return utils.GetRequestHandler(ctx, c.GetSubmissionInfo(), utils.RequestBody{Data: query})
	})

	c.FiberApp.Post("/quiz/mint", func(ctx *fiber.Ctx) error {
		body := new(NFTClaimAttempt)
		return utils.PostRequestHandler(ctx, c.MintQuizNFT(), utils.RequestBody{Data: body})
	})
}
