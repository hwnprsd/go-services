package models

import (
	"time"

	"gorm.io/gorm"
)

type QuizSubmission struct {
	ID                   uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt            time.Time      `json:"created_at,omitempty"`
	UpdatedAt            time.Time      `json:"updated_at,omitempty"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Email                *string        `json:"email,omitempty"`
	QuizID               string         `json:"quiz_id,omitempty"`
	ClaimID              string         `json:"claim_id,omitempty" gorm:"uniqueIndex"`
	IsNFTClaimed         bool           `json:"is_claimed"`
	NFTClaimAttemptCount uint           `json:"nft_claim_attempt_count"`
	IsNFTClaimMailSent   bool           `json:"is_nft_claim_mail_sent"`
}
