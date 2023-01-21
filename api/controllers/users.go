package controllers

import (
	"encoding/json"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	"github.com/gofiber/fiber/v2"
)

type CreateUserBody struct {
	Email string `json:"email,omitempty" validate:"required,email,min=6,max=32"`
	Level uint   `json:"level"`
}

// Add users with unique emails to the table
func (ctrl *Controller) CreateUser() func(utils.RequestBody) interface{} {
	return func(data utils.RequestBody) interface{} {
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
			panic(result.Error)
		}

		jsonString1, _ := json.Marshal(fiber.Map{
			"email": body.Email,
			"type":  "WELCOME",
		})

		jsonString2, _ := json.Marshal(fiber.Map{
			"email": body.Email,
			"to":    "scheduler",
		})

		ctrl.MQ.MailerQueue.PublishMessage(string(jsonString1))
		ctrl.MQ.SchedulerQueue.PublishMessage(string(jsonString2))

		return "User Added"
	}
}

func (ctrl *Controller) GetUsers() func(utils.RequestBody) interface{} {
	return func(data utils.RequestBody) interface{} {
		users := []models.User{}
		result := ctrl.DB.Find(&users)
		if result.Error != nil {
			panic(result.Error)
		}
		return users
	}
}
