package models

type Tag struct {
	ID       uint      `gorm:"primarykey" json:"id,omitempty"`
	Tag      string    `json:"tag"`
	Articles []Article `json:"articles" gorm:"many2many:articles_tags;"`
	RssFeeds []RssFeed `json:"rss_feeds" gorm:"many2many:rss_tags;"`
}
