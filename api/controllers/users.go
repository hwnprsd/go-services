package controllers

import (
	"log"
	"net/http"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	fiberw "github.com/hwnprsd/go-easy-docs/fiber-wrapper"
	"github.com/hwnprsd/shared_types"
	"gorm.io/gorm/clause"
)

func (c *Controller) SetupUserRoutes() {
	group := fiberw.NewGroup(c.FiberApp, "/users")
	fiberw.Get(group, "/", c.GetUsers)
	fiberw.Post(group, "/create", c.SubscribeWeb3NewLetter)

}

func (ctrl *Controller) getUser(uniqueId string) (*models.User, error) {
	user := models.User{}
	resp := ctrl.DB.Where("unique_id = ?", uniqueId).First(&user)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &user, nil
}

func (ctrl *Controller) createUser(data *models.User) (*models.User, error) {
	resp := ctrl.DB.Create(data).Clauses(clause.Returning{})
	if resp.Error != nil {
		return nil, resp.Error
	}
	return data, nil
}

func (ctrl *Controller) GetOrCreateUser(data *models.User) (*models.User, error) {
	user, err := ctrl.getUser(data.UniqueId)
	if err != nil {
		log.Println("Error fetching user with ID -", data.UniqueId)
		log.Println(err)
		log.Println("Creating user")
		log.Println("---------------------")
		user, err = ctrl.createUser(data)
		if err != nil {
			log.Println("Error creating user with ID -", data.UniqueId)
			log.Println(err)
			log.Println("---------------------")
			return nil, err
		}
	}
	return user, nil
}

type SubscribeWeb3NewsletterBody struct {
	Email string `json:"email" validate:"required,email,min=6,max=32"`
}

// Add users with unique emails to the table
func (ctrl *Controller) SubscribeWeb3NewLetter(body SubscribeWeb3NewsletterBody) (*string, error) {
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
