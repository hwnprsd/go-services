package gpt

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"flaq.club/workers/utils"
	"github.com/hwnprsd/shared_types"
	"github.com/sashabaranov/go-openai"
	"github.com/streadway/amqp"
)

type GptHandler struct {
	ApiQueue *utils.Queue
}

func NewGptHandler(apiQueue *utils.Queue) *GptHandler {
	return &GptHandler{apiQueue}
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
		// Enable when live
		h.SummarizeBlog(message)
		break
	}
	// Just ack everything
	payload.Ack(false)
}

func summarize(client *openai.Client, messages []openai.ChatCompletionMessage) (*string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)
	if err != nil {
		return nil, err
	}
	return &resp.Choices[0].Message.Content, nil
}

func (h *GptHandler) SummarizeBlog(data shared_types.SummarizeBlogMessage) {
	h.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
		data.TaskId,
		"SUMMARY_START",
		"",
	))
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(OPENAI_API_KEY)

	paragraphs := (strings.Split(data.ScrapeData, "\n"))

	log.Println(paragraphs)
	log.Println("---------")

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
		Content: "You are a helpful text summrizing bot. Your job is to summarize of the following content into a short 800 word summary",
	}
	messages = append(messages, contextMessage)

	recentSummary := ""
	for _, b := range batches {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: b,
		})
		// Modify the messages
		summary, err := summarize(client, messages)
		if err != nil {
			log.Println("Error summarizing")
			log.Println(err)
			return
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

	finalSummary := recentSummary

	log.Println("Summary - ", finalSummary)

	h.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
		data.TaskId,
		"SUMMARY_COMPLETE",
		finalSummary,
	))
}
