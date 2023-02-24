package models

import (
	"time"

	"gorm.io/gorm"
)

// eventName := "Flaq x Web3 Gorkha Siliguri"
// eventDescription := "Thank you for attending Flaq x Web3 Gorkha Siliguri Mixer. This POAP is a proof that you attended the event."
// userName := "Ashwin Prasad"
// eventKey := "event-2"
// fontSize := 17
// textX := 47
// textY := 532
// bucketName := "flaq-nfts"

// TODO Make this more dynamic in terms of handling dynamic templates
type Web3Event struct {
	ID               uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt        time.Time      `json:"created_at,omitempty"`
	UpdatedAt        time.Time      `json:"updated_at,omitempty"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	EventName        string         `json:"event_name"`
	EventDescription string         `json:"event_description"`
	EventKey         string         `json:"event_key"`
	FontSize         uint           `json:"font_size"`
	TextX            uint           `json:"text_x"`
	TextY            uint           `json:"text_y"`
	BucketName       string         `json:"bucket_name"`
}
