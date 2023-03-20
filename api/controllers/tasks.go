package controllers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	"github.com/gofiber/fiber/v2"
	fiberw "github.com/hwnprsd/go-easy-docs/fiber-wrapper"
	"github.com/hwnprsd/shared_types"
)

func (c *Controller) SetupTaskRoutes() {
	group := fiberw.NewGroup(c.FiberApp, "/tasks")

	fiberw.GetWithExtra(group, "/", c.GetTaskDetails, func(ctx *fiber.Ctx) (string, error) {
		taskId := ctx.Query("taskId")
		return taskId, nil
	})

	fiberw.Post(group, "/schedule/email", c.ScheduleEmail)
}

type GetTaskDetailsQuery struct {
	TaskID string `json:"task_id"`
}

func (ctrl *Controller) GetTaskDetails(taskID string) (*models.Task, error) {
	if taskID == "" {
		return nil, &utils.RequestError{StatusCode: 400, Message: "Error fetching task data", Err: errors.New("Error finding task for id")}
	}
	task := models.Task{}
	res := ctrl.DB.Where("id = ?", taskID).First(&task)
	if res.Error != nil {
		return nil, &utils.RequestError{StatusCode: 400, Message: "Error fetching task data", Err: res.Error}
	}

	return &task, nil
}

type ScheduleEmailBody struct {
	CampaignId uint `json:"campaign_id" validate:"required"`
	MinsLater  uint `json:"mins_later" validate:"required"`
}

func (ctrl *Controller) ScheduleEmail(body ScheduleEmailBody) (*string, error) {
	scheduleTime := time.Now().Add(time.Minute * time.Duration(body.MinsLater))
	// scheduleTime := time.Now()
	log.Println(time.Now().String())
	log.Println(scheduleTime.String())
	message := shared_types.NewScheduleEmailMessage(0, body.CampaignId, scheduleTime, map[string]string{"test": "true"})
	error := ctrl.MQ.SchedulerQueue.PublishMessage(message)

	if error != nil {
		log.Fatal(error)
	}

	response := fmt.Sprintf("task scheduled for job ID %d", body.CampaignId)
	return &response, nil
}
