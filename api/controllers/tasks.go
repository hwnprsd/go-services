package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"flaq.club/api/models"
	"flaq.club/api/utils"
	"github.com/gofiber/fiber/v2"
	fiberw "github.com/hwnprsd/go-easy-docs/fiber-wrapper"
	"github.com/hwnprsd/shared_types"
	"gorm.io/gorm/clause"
)

func (c *Controller) SetupTaskRoutes() {
	group := fiberw.NewGroup(c.FiberApp, "/tasks")

	fiberw.GetWithExtra(group, "/", c.GetTaskDetails, func(ctx *fiber.Ctx) (string, error) {
		taskId := ctx.Query("taskId")
		return taskId, nil
	}).WithQuery("taskId")

	fiberw.Post(group, "/schedule/email", c.ScheduleEmail)
	fiberw.PostWithExtra(group, "/summarize", c.SummarizeBlog, c.InjectUser)
	fiberw.Get(group, "/update/rss", c.UpdateRssFeed)
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

type SummarizeBlog struct {
	Url       string `json:"url"`
	WordCount uint   `json:"word_count"`
}

func (ctrl *Controller) SummarizeBlog(data SummarizeBlog, user *models.User) (*string, error) {
	log.Println("Dealing with user -", *user.Email)
	SCRAPER_TOKEN := os.Getenv("SCRAPER_TOKEN")
	scraperUrl := ("https://wln0664peg.execute-api.ap-south-1.amazonaws.com/default/jsScraper3")
	body, _ := json.Marshal(utils.Map{
		"url": data.Url,
	})
	req, err := http.NewRequest("POST", scraperUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Error creating request")
	}
	req.Header.Set("x-api-key", SCRAPER_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error creating request")
	}
	bodyBytes, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	task := models.Task{
		Category: "BLOG_SUMMARY",
		Status:   "STARTED",
		UserID:   user.ID,
	}

	if r := ctrl.DB.Clauses(clause.Returning{}).Create(&task); r.Error != nil {
		return nil, fiberw.NewRequestError(400, "Error creating a task", r.Error)
	}
	//
	type ScraperResponse struct {
		TextContent string `json:"textContent"`
	}

	responseBody := ScraperResponse{}
	json.Unmarshal(bodyBytes, &responseBody)

	scrapedData := responseBody.TextContent

	log.Println("Scraped Data", scrapedData)

	ctrl.MQ.GPTQueue.PublishMessage(shared_types.NewSummarizeBlogMessage(
		task.ID,
		scrapedData,
	))

	// TODO: Work on a credits system

	returnable := fmt.Sprintf("%d", task.ID)
	return &returnable, nil
}

func (ctrl Controller) UpdateRssFeed() (*string, error) {
	message := shared_types.NewUpdateRssData(1)
	ctrl.MQ.GPTQueue.PublishMessage(message)
	response := "RSS Update Queued"
	return &response, nil
}
