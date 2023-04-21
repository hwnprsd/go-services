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
	fiberw.Get(group, "/newsletter/update/rss", c.UpdateRssFeed)
	fiberw.Get(group, "/newsletter/summary/create", c.CreateRssNewsletter)
	fiberw.Post(group, "/newsletter/summary/send", c.SendRssNewsletter)
	fiberw.PostWithExtra(group, "/cv/analyse", c.ParsePdfCv, c.InjectUser)

	fiberw.RawPost(group, "/cv/analyse/upload/sync", c.ParsePdfCvUploadSync)
	fiberw.Post(group, "/cv/analyse/url/sync", c.ParsePdfCvUrlSync)
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

func (ctrl Controller) CreateRssNewsletter() (*string, error) {
	now := time.Now()
	loc, _ := time.LoadLocation("Local")
	publishedDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	message := shared_types.NewCreateRssNewsletterMessage(1, "Technology", publishedDate)
	ctrl.MQ.GPTQueue.PublishMessage(message)
	response := "RSS Update Queued"
	return &response, nil
}

type SendRssNewsletterBody struct {
	Tag               string    `json:"tag"`
	Date              time.Time `json:"date"`
	ScheduleMinsLater uint      `json:"schedule_mins_later"`
	CampaignId        uint      `json:"campaign_id"`
}

func (ctrl Controller) SendRssNewsletter(body SendRssNewsletterBody) (*string, error) {
	y, m, d := time.Now().Date()
	publishedDate := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	scheduleTime := time.Now().Add(time.Minute * time.Duration(body.ScheduleMinsLater))
	message := shared_types.NewSendRssNewsletterMessage(1, body.CampaignId, body.Tag, publishedDate, scheduleTime)
	ctrl.MQ.MailerQueue.PublishMessage(message)
	response := "RSS Send Email Message Sent"
	return &response, nil
}

type ParsePdfCvBody struct {
	Url string `json:"url"`
}

func (ctrl Controller) ParsePdfCv(data ParsePdfCvBody, user *models.User) (*string, error) {
	task := models.Task{
		Category: "PDF_CV_ANALYSIS",
		Status:   "PDF_EXTRACT_STARTED",
		UserID:   user.ID,
	}

	if r := ctrl.DB.Clauses(clause.Returning{}).Create(&task); r.Error != nil {
		return nil, fiberw.NewRequestError(400, "Error creating a task", r.Error)
	}
	ctrl.MQ.GPTQueue.PublishMessage(shared_types.NewPdfParseMessage(task.ID, data.Url))
	taskId := fmt.Sprint(task.ID)
	return &taskId, nil
}

type ParsePdfSyncResponse struct {
	Success    bool   `json:"success"`
	TaskId     uint   `json:"task_id"`
	Error      string `json:"error,omitempty"`
	ErrorCode  uint   `json:"error_code,omitempty"`
	ParsedData any    `json:"parsed_data"`
}

func (ctrl Controller) handlePdfAnalysis(t *models.Task) *ParsePdfSyncResponse {
	for i := 0; i < 50; i++ {
		t, _ := ctrl.GetTaskDetails(fmt.Sprint(t.ID))
		if t.Status == "CV_EXTRACTION_STARTED" || t.Status == "CV_EXTRACTION_COMPLETE" {
			time.Sleep(5 * time.Second)
			continue
		} else if t.Status == "CV_EXTRACTION_FAILED_1" {
			return &ParsePdfSyncResponse{
				Success:   false,
				TaskId:    t.ID,
				Error:     t.Data,
				ErrorCode: 1,
			}
		} else if t.Status == "CV_EXTRACTION_FAILED_2" {
			return &ParsePdfSyncResponse{
				Success:   false,
				TaskId:    t.ID,
				Error:     t.Data,
				ErrorCode: 2,
			}
		} else if t.Status == "CV_ANALYSIS_FAILED" {
			return &ParsePdfSyncResponse{
				Success:   false,
				TaskId:    t.ID,
				Error:     t.Data,
				ErrorCode: 3,
			}
		}
		var data interface{}
		err := json.Unmarshal([]byte(t.Data), &data)
		if err != nil {
			return &ParsePdfSyncResponse{
				ParsedData: t.Data,
				Success:    false,
				TaskId:     t.ID,
				Error:      "Error parsing json. Try again",
				ErrorCode:  100,
			}
		}
		return &ParsePdfSyncResponse{
			ParsedData: data,
			Success:    true,
			TaskId:     t.ID,
		}
	}

	return &ParsePdfSyncResponse{
		Success:   false,
		Error:     "Timed Out",
		ErrorCode: 4,
		TaskId:    t.ID,
	}
}

func (ctrl Controller) ParsePdfCvUrlSync(data ParsePdfCvBody) (*ParsePdfSyncResponse, error) {
	task := models.Task{
		Category: "PDF_CV_ANALYSIS",
		Status:   "CV_EXTRACTION_STARTED",
		UserID:   9,
	}

	if r := ctrl.DB.Clauses(clause.Returning{}).Create(&task); r.Error != nil {
		return nil, fiberw.NewRequestError(400, "Error creating a task", r.Error)
	}
	ctrl.MQ.GPTQueue.PublishMessage(shared_types.NewPdfParseMessage(task.ID, data.Url))
	return ctrl.handlePdfAnalysis(&task), nil
}

func (ctrl Controller) ParsePdfCvUploadSync(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(400).SendString(fmt.Sprint(utils.Map{
			"success": false,
			"error":   "Error extracting PDF from form-data. Make sure you send header `content-type: multipart/form-data` with a form filed called `file` with the PDF'",
		}))
	}
	file, _ := fileHeader.Open()
	fileBytes := make([]byte, fileHeader.Size)
	_, _ = file.Read(fileBytes)
	task := models.Task{
		Category: "PDF_CV_ANALYSIS_BYTES",
		Status:   "CV_EXTRACTION_STARTED",
		UserID:   9,
	}

	if r := ctrl.DB.Clauses(clause.Returning{}).Create(&task); r.Error != nil {
		return ctx.Status(400).JSON(utils.Map{
			"success": false,
			"error":   "Error creating a task for user. Contact adming",
		})
	}
	ctrl.MQ.GPTQueue.PublishMessage(shared_types.NewPdfParseBytesMessage(task.ID, fileBytes))
	res := ctrl.handlePdfAnalysis(&task)
	return ctx.Status(200).JSON(*res)
}
