package models

import (
	"time"

	"gorm.io/gorm"
)

type NftMint struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	TxnID     string         `json:"txn_id,omitempty"`
	Address   string         `json:"address,omitempty"`
}
