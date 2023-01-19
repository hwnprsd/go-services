package controllers

import (
	"flaq.club/api/models"
	"flaq.club/api/utils"
)

type CreateUserBody struct {
	Email string `json:"email,omitempty"`
}

// Add users with unique emails to the table
func (ctrl *Controller) CreateUser() func(utils.RequestBody) interface{} {
	return func(data utils.RequestBody) interface{} {
		body := data.Data.(*CreateUserBody)
		user := models.User{
			Email: &body.Email,
		}
		result := ctrl.DB.Create(&user)
		if result.Error != nil {
			panic(result.Error)
		}

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
