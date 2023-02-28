package controllers

import (
	"net/http"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/hwnprsd/shared_types"
)

func (c *Controller) SetupUserRoutes() {
	group := c.FiberApp.Group("users")
	group.Post("/create", func(ct *fiber.Ctx) error {
		body := new(CreateUserBody)
		return utils.PostRequestHandler(ct, c.CreateUser(), utils.RequestBody{Data: body})
	})

	group.Get("/", func(ctx *fiber.Ctx) error {
		return utils.GetRequestHandler(ctx, c.GetUsers(), utils.RequestBody{})
	})
}

type CreateUserBody struct {
	Email string `json:"email,omitempty" validate:"required,email,min=6,max=32"`
	Level uint   `json:"level"`
}

// Add users with unique emails to the table
func (ctrl *Controller) CreateUser() utils.PostHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		body := data.Data.(*CreateUserBody)
		if body.Level < 1 {
			body.Level = 1
		}
		user := models.User{
			Email: &body.Email,
			Level: body.Level,
		}
		result := ctrl.DB.Create(&user)
		if result.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error Creating User in the Database",
				Err:        result.Error,
			}
		}

		ctrl.MQ.MailerQueue.PublishMessage(shared_types.NewSendMailMessage(
			0,
			"ashwin@solace.money",
			"Claim sthe NFT you jusat earned! Again 2 - ! 3x",
			1,
			map[string]string{"image_url": "https://flaq-assets.s3.ap-south-1.amazonaws.com/backgrounds/TTWE.png",
				"title":       "The term 'Web3' explained!",
				"body":        "While the features of the new web are endless, the main features, as we see it, are that web3 is permissionless, decentralized, trustless and money will become a native feature of the internet! All of this together will not just lead to ownership, but also anti-censoring. With web3 being the future, why not learn all about it - step by step? Find a more elaborate version of this summary, and learn more about what the term ‘web3’ means here ",
				"button_link": "https://bit.ly/flaq-term-web3-explained",
				"button_text": "Learn More",
			},
		))
		ctrl.MQ.SchedulerQueue.PublishMessage(utils.Map{
			"type": "WELCOME_2_SCHEDULE",
			"data": utils.Map{
				"email": body.Email,
			},
		})
		ctrl.MQ.NftQueue.PublishMessage(utils.Map{
			"type": "NFT_SCHEDULE",
			"data": utils.Map{
				"email": body.Email,
			},
		})

		return "User Added", nil
	}
}

func (ctrl *Controller) GetUsers() utils.GetHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		users := []models.User{}
		result := ctrl.DB.Find(&users)
		if result.Error != nil {
			return nil, &utils.RequestError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error fetching data",
				Err:        result.Error,
			}
		}
		return users, nil
	}
}
