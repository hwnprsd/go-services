package gpt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"flaq.club/workers/utils"
	"github.com/hwnprsd/shared_types"
	"github.com/sashabaranov/go-openai"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type GptHandler struct {
	ApiQueue *utils.Queue
	Db       *gorm.DB
}

func NewGptHandler(apiQueue *utils.Queue, db *gorm.DB) *GptHandler {
	return &GptHandler{apiQueue, db}
}

func (h *GptHandler) HandleMessages(payload *amqp.Delivery) {
	baseMessage := shared_types.MessagingBase{}
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			payload.Reject(false)
		}
	}()
	err := json.Unmarshal(payload.Body, &baseMessage)
	if err != nil {
		log.Printf("Error parsing JSON message. Please check what the sender sent! QUEUE - %s", payload.Body)
		payload.Reject(false)
		return
	}
	switch baseMessage.WorkType {
	case shared_types.WORK_TYPE_SUMMARIZE_BLOG:
		message := shared_types.SummarizeBlogMessage{}
		json.Unmarshal(payload.Body, &message)
		log.Println("Summarizing Blog")
		err := h.SummarizeBlog(message)
		if err != nil {
			log.Println("--- Error summarizing blog")
			log.Println(err)
		}
		break
	case shared_types.WORK_TYPE_UPDATE_RSS:
		message := shared_types.UpdateRssData{}
		json.Unmarshal(payload.Body, &message)
		err := h.ReadRSSFeed()
		if err != nil {
			log.Println("--- Error updating RSS Feed", err)
		}
		break
	case shared_types.WORK_TYPE_CREATE_RSS_NEWSLETTER:
		message := shared_types.CreateRssNewsletterMessage{}
		json.Unmarshal(payload.Body, &message)
		err := h.CreateAndSendNewsletter(message)
		if err != nil {
			log.Println("--- Error creating summary newsletter", err)
		}
		break

	case shared_types.WORK_TYPE_PDF_PARSE_CV:
		message := shared_types.PdfParseCVMessage{}
		json.Unmarshal(payload.Body, &message)
		err := h.ParsePdfCv(message)
		if err != nil {
			log.Println("--- Error extacting or anaylising PDF CV", err)
		}
		break
	case shared_types.WORK_TYPE_PDF_PARSE_CV_BYTES:
		message := shared_types.PdfParseCVBytesMessage{}
		json.Unmarshal(payload.Body, &message)
		err := h.ParsePdfCvBytes(message)
		if err != nil {
			log.Println("--- Error extacting or anaylising PDF CV using Bytes", err)
		}
		break

	}

	// Just ack everything
	payload.Ack(false)
}

func summarizeBlog(blogContent string, wordCount uint) (*string, error) {
	paragraphs := (strings.Split(blogContent, "\n"))

	gist := blogContent[:10]
	log.Println("Summarizing", gist, "...")

	batches := make([]string, 0)

	batch := ""

	// go from paragraph to paragraph - add about 1000 characters to a new buffer
	charCount := 0
	for i := 0; i < len(paragraphs); i++ {
		current := paragraphs[i]
		batch += current
		charCount += len(current)
		if charCount > 2000 {
			// Append the batch and reset it
			batches = append(batches, batch)
			batch = ""
			charCount = 0
		}
	}

	messages := make([]openai.ChatCompletionMessage, 0)
	contextMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: fmt.Sprintf("You are a helpful text summrizing bot. Your job is to summarize of the following content into a short %d word summary", wordCount),
	}
	messages = append(messages, contextMessage)

	recentSummary := ""
	for _, b := range batches {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: b,
		})
		// Modify the messages
		summary, err := GetCompletion(messages)
		if err != nil {
			log.Println("Error summarizing")
			log.Println(err)
			return nil, errors.New(fmt.Sprintf("Error summarizing - %e", err))
		}
		// summary := fmt.Sprintf("SUMMARY %d", i)
		recentSummary = *summary
		// Remove the original message which created the response
		messages = messages[:len(messages)-1]
		// Add the newly made summary
		messages = make([]openai.ChatCompletionMessage, 0)
		messages = append(messages, contextMessage)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: *summary,
		})
	}

	log.Println("Summary complete -", gist)

	return &recentSummary, nil

}

// Summarize given article summaries into one newsletter summary
func SummarizeSummaries(summaries string) (*string, error) {
	log.Println("Summarizing summaries")
	messages := make([]openai.ChatCompletionMessage, 0)
	contextMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: fmt.Sprintf("You are a helpful text summrizing bot. You should summarize the given summaries in an unordered list"),
	}
	messages = append(messages, contextMessage)
	summaryMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: summaries,
	}
	messages = append(messages, summaryMessage)
	ret, err := GetCompletion(messages)
	log.Println("Summarizing Complete")
	return ret, err
}

func GetCvAnalysis(cvText string) (*string, error) {
	log.Println("Getting information on the CV")
	messages := make([]openai.ChatCompletionMessage, 0)
	contextMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: fmt.Sprint("You are a helpful CV / Resume analyzer bot. You will be given a CV from a candidate's PDF. You have to return all the information you find strictly in this JSON format { \"name\": \"\", \"email\": \"\", \"dob\": \"\", \"technicalSkills\": [\"\"], \"education\": [{\"institution\": \"\", \"degree\": \"\", \"field\": \"\", \"graduationDate\": \"\"}], \"workExperience\": [ {\"companyName\": \"\", \"duration\": \"\", \"startDate\": \"\", \"position\": \"\", \"location\": \"\", \"responsibilities\": \"\"} ], \"awards\": [\"\"] }"),
	}
	messages = append(messages, contextMessage)
	cvMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: cvText,
	}
	messages = append(messages, cvMessage)
	ret, err := GetCompletion(messages)
	log.Println("CV Analysis complete")
	if err != nil {
		log.Println(err)
	}
	// TODO: Check if the CV is scanned properly
	return ret, err
}

// Given a slice of messages, call the openai API and return the response
func GetCompletion(messages []openai.ChatCompletionMessage) (*string, error) {
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(OPENAI_API_KEY)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)
	if err != nil {
		log.Println("OpenAI Error")
		// TODO: Check if this is an API error. If so, retry
		return nil, err
	}
	return &resp.Choices[0].Message.Content, nil
}

func (h *GptHandler) SummarizeBlog(data shared_types.SummarizeBlogMessage) error {
	h.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
		data.TaskId,
		"SUMMARY_START",
		"",
	))
	summary, err := summarizeBlog(data.ScrapeData, 800)
	log.Println("Summary - ", summary)

	h.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
		data.TaskId,
		"SUMMARY_COMPLETE",
		summary,
	))
	return err
}
