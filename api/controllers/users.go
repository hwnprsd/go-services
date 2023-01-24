package controllers

import (
	"net/http"

	"flaq.club/api/models"
	"flaq.club/api/utils"
)

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

		ctrl.MQ.MailerQueue.PublishMessage(utils.Map{
			"type": "WELCOME",
			"data": utils.Map{
				"email": body.Email,
			},
		})
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
