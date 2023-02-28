package controllers

import (
	"errors"
	"log"
	"time"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/hwnprsd/shared_types"
)

func (c *Controller) SetupTaskRoutes() {
	group := c.FiberApp.Group("tasks")

	group.Get("/", func(ctx *fiber.Ctx) error {
		taskId := ctx.Query("taskId")
		query := new(GetTaskDetailsQuery)
		query.TaskID = taskId
		return utils.GetRequestHandler(ctx, c.GetTaskDetails(), utils.RequestBody{Query: query})
	})

	group.Post("/schedule/email", func(ctx *fiber.Ctx) error {
		body := ScheduleEmailBody{}
		return utils.PostRequestHandler(ctx, c.ScheduleEmail(), utils.RequestBody{Data: &body})
	})
}

type GetTaskDetailsQuery struct {
	TaskID string `json:"task_id"`
}

func (ctrl *Controller) GetTaskDetails() utils.GetHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		query := data.Query.(*GetTaskDetailsQuery)
		if query.TaskID == "" {
			return nil, &utils.RequestError{StatusCode: 400, Message: "Error fetching task data", Err: errors.New("Error finding task for id")}
		}
		task := models.Task{}
		res := ctrl.DB.Where("id = ?", query.TaskID).First(&task)
		if res.Error != nil {
			return nil, &utils.RequestError{StatusCode: 400, Message: "Error fetching task data", Err: res.Error}
		}

		return task, nil
	}
}

type ScheduleEmailBody struct {
	ListName string `json:"list_name"`
}

func (ctrl *Controller) ScheduleEmail() utils.PostHandler {
	return func(data utils.RequestBody) (interface{}, error) {
		scheduleTime := time.Now().Add(time.Minute * time.Duration(1))
		// scheduleTime := time.Now()
		log.Println(time.Now().String())
		log.Println(scheduleTime.String())
		message := shared_types.NewScheduleEmailMessage(0, 1, scheduleTime, map[string]string{"test": "true"})
		error := ctrl.MQ.SchedulerQueue.PublishMessage(message)
		message = shared_types.NewScheduleEmailMessage(0, 2, scheduleTime, map[string]string{"test": "true"})
		error = ctrl.MQ.SchedulerQueue.PublishMessage(message)
		message = shared_types.NewScheduleEmailMessage(0, 3, scheduleTime, map[string]string{"test": "true"})
		error = ctrl.MQ.SchedulerQueue.PublishMessage(message)

		if error != nil {
			log.Fatal(error)
		}
		return "task scheduled for job ID 0", nil
	}
}
