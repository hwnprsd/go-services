package scraper

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"flaq.club/workers/utils"
	"github.com/chromedp/chromedp"
	"github.com/hwnprsd/shared_types"
	"github.com/streadway/amqp"
	"golang.org/x/net/html"
)

type ScraperHandler struct {
	ApiQueue *utils.Queue
}

func NewScraperHandler(apiQueue *utils.Queue) *ScraperHandler {
	return &ScraperHandler{apiQueue}
}

func (h *ScraperHandler) HandleMessages(payload *amqp.Delivery) {
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
	case shared_types.WORK_TYPE_SCRAPE_URL:
		message := shared_types.ScrapeUrlMessage{}
		json.Unmarshal(payload.Body, &message)
		log.Println("Asking to mint POAP when disabled")
		// Enable when live
		h.ScrapeUrl(message)
		break
	}
}

func (h *ScraperHandler) ScrapeUrl(message shared_types.ScrapeUrlMessage) {
	useChrome(message.Url)
	// resp, err := http.Get(message.Url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	//
	// doc, err := html.Parse(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// text := extractText(doc)
	// log.Println(text)
}

func extractText(n *html.Node) string {
	var sb strings.Builder
	if n.Type == html.TextNode && n.Data != "" {
		sb.WriteString(n.Data)
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return ""
	}

	if n.Type == html.ElementNode && (n.Data == "h1" || n.Data == "h2" || n.Data == "h3" || n.Data == "p" || n.Data == "code") {
		sb.WriteString(" ")
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(extractText(c))
	}

	return sb.String()
}

func useChrome(url string) {
	// Create a new context
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), getDebugURL())
	defer cancel()

	// Set a timeout
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Navigate to the target webpage
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the React component to render and load the data
	err = chromedp.Run(ctx, chromedp.WaitVisible(".react-component", chromedp.ByQuery))
	if err != nil {
		log.Fatal(err)
	}

	// Extract the text content of the React component
	var text string
	err = chromedp.Run(ctx, chromedp.Text(".react-component", &text, chromedp.ByQuery))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(text)

}

func getDebugURL() string {
	resp, err := http.Get(os.Getenv("CHROME_DP_URL"))
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result["webSocketDebuggerUrl"].(string)
}
