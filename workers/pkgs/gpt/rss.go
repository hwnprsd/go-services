package gpt

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"flaq.club/workers/models"
	"github.com/hwnprsd/shared_types"
	"github.com/mmcdole/gofeed"
)

func (handler *GptHandler) CreateAndSendNewsletter(data shared_types.CreateRssNewsletterMessage) error {
	// 1. Get all the summaries for a particular date and tag
	// 2. Further summarize them to create the newsletter content
	// 3. Draft a newsletter with the relevant links and the summary content appropriately
	// 5. Email it out

	newsletterData, err := handler.createNewsletter(data.Tag, data.Date)
	if err != nil {
		log.Println("Error Creating Summary Newsletter", err)
		return err
	}

	// store newsletter
	articles := make([]models.Article, 0)
	for _, article := range newsletterData.Articles {
		articles = append(articles, models.Article{
			ID: article.ID,
		})
	}
	newsletter := models.SummaryNewsletter{
		Summary:  newsletterData.Summary,
		Articles: articles,
	}

	if resp := handler.Db.Create(&newsletter); resp.Error != nil {
		log.Println("Error creating a newsletter", err)
		return err
	}

	return nil
}

type Article struct {
	Title string
	Url   string
	ID    uint
}

type NewsletterSummmary struct {
	Summary  string
	Articles []Article
}

func (handler *GptHandler) createNewsletter(tagName string, date time.Time) (*NewsletterSummmary, error) {
	// 1. Fetch all articles by Tag Name
	// 2. Batch them into groups of 5
	// 3. Append summaries
	tag := models.Tag{}
	if res := handler.Db.Preload("Articles").Where("tag = ?", tagName).First(&tag); res.Error != nil {
		log.Println(res.Error)
		return nil, res.Error
	}

	summaries := make([]string, 0)
	articles := make([]Article, 0)
	batchCount := 7
	count := 0
	totalCount := 0
	summaryConcat := ""

	for _, article := range tag.Articles {
		// Only care for the required date
		if !article.PublishDate.Equal(date) {
			continue
		}
		totalCount++
		count++
		if count >= batchCount {
			summaries = append(summaries, summaryConcat)
			summaryConcat = ""
			count = 0
		}
		summaryConcat += article.Summary + "\n\n"
		articles = append(articles, Article{article.Title, article.Url, article.ID})
	}

	log.Println(fmt.Printf("Summarizing %d articles summaries", totalCount))

	if len(summaries) == 0 {
		log.Println("Nothing to summarize")
		return nil, nil
	}

	finalSummary := ""
	for _, summaryConcat = range summaries {
		summary, err := SummarizeSummaries(summaryConcat)
		if err != nil {
			log.Println("Error summarizing the summaries", err)
			return nil, err
		}
		finalSummary += "\n" + *summary

	}

	data := NewsletterSummmary{
		Summary:  finalSummary,
		Articles: articles,
	}

	return &data, nil
}

func (handler *GptHandler) ReadRSSFeed() error {
	// 1. Get all RSS feed URLs
	// 2. Fetch blogs from these RSS feeds
	// 3. Check if they already exist - If yes - continue
	// 4. If no, summarize and store it in the article databas

	// Get all RSS Feeds
	rss := []models.RssFeed{}
	if resp := handler.Db.Where("should_skip = ?", false).Preload("Tags").Find(&rss); resp.Error != nil {
		return errors.New(fmt.Sprintf("1. Error fetching RSS Feeds, %e", resp.Error))
	}

	if len(rss) == 0 {
		return errors.New("2. No RSS Feeds in the database")
	}

	for _, rssFeed := range rss {
		// Fetch blogs
		feed, err := fetchBlogs(rssFeed.Url)
		if err != nil {
			log.Println("Error fetching blogs for url ", rssFeed.Url, err)
			continue
		}
		blogs := feed.Items
		for _, blog := range blogs {
			// Check if the blog is already in the database
			count := int64(0)
			handler.Db.Model(models.Article{}).Where("guid = ?", blog.GUID).Count(&count)

			if count != 0 {
				log.Println("Blog already exists with GUID", blog.GUID)
				continue
			}

			summary, err := summarizeBlog(blog.Content, 200)
			if err != nil {
				log.Println("Error summarizing blog")
				log.Println(err)
				continue
			}

			authors := make([]string, 0)
			tags := make([]models.Tag, 0)

			for _, author := range blog.Authors {
				authors = append(authors, author.Name)
			}

			for _, tag := range rssFeed.Tags {
				tags = append(tags, models.Tag{
					ID: tag.ID,
				})
			}

			var year, day int
			var month time.Month
			loc, _ := time.LoadLocation("Local")
			if blog.PublishedParsed != nil {
				year, month, day = blog.PublishedParsed.Date()
				loc = blog.PublishedParsed.Location()
			}

			publishedDate := time.Date(year, month, day, 0, 0, 0, 0, loc)
			// Create article and store summary
			article := models.Article{
				Summary:     *summary,
				Title:       blog.Title,
				Authors:     strings.Join(authors, ", "),
				PublishDate: publishedDate,
				Url:         blog.Link,
				Tag:         tags,
				GUID:        blog.GUID,
			}

			if resp := handler.Db.Create(&article); resp.Error != nil {
				log.Println("Error creating an article")
				continue
			}
		}
	}
	log.Println("Summarization and storage complete")
	return nil
}

func fetchBlogs(url string) (*gofeed.Feed, error) {
	feedParser := gofeed.NewParser()
	return feedParser.ParseURL(url)
}
