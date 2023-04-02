package controllers

import (
	"net/http"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	fiberw "github.com/hwnprsd/go-easy-docs/fiber-wrapper"
	"github.com/hwnprsd/shared_types"
)

func (c *Controller) SetupUserRoutes() {
	group := fiberw.NewGroup(c.FiberApp, "/users")
	fiberw.Get(group, "/", c.GetUsers)
	fiberw.Post(group, "/create", c.CreateUser)

}

type CreateUserBody struct {
	Email string `json:"email" validate:"required,email,min=6,max=32"`
}

// Add users with unique emails to the table
func (ctrl *Controller) CreateUser(body CreateUserBody) (*string, error) {
	user := models.User{
		Email: &body.Email,
	}
	result := ctrl.DB.Create(&user)
	if result.Error != nil {
		return nil, &fiberw.RequestError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error Creating User in the Database",
			Err:        result.Error,
		}
	}

	// Subscribe newsletter
	ctrl.MQ.MailerQueue.PublishMessage(shared_types.NewAddUserToListMessage(
		2,
		body.Email,
		"Flaq User",
		"newsletter",
	))

	welcomeEmailTemplateId := 8
	// Send welcome email
	ctrl.MQ.MailerQueue.PublishMessage(shared_types.NewSendMailMessage(
		2,                            // task ID
		body.Email,                   // Email
		"[Flaq] Welcome to Flaq",     // Subject?
		uint(welcomeEmailTemplateId), // Template Id
		nil,                          // Data Map
	))

	resp := "User Added"
	return &resp, nil
}

func (ctrl *Controller) GetUsers() (*[]models.User, error) {
	users := []models.User{}
	result := ctrl.DB.Find(&users)
	if result.Error != nil {
		return nil, &utils.RequestError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error fetching data",
			Err:        result.Error,
		}
	}
	return &users, nil
}
