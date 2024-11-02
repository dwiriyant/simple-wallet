package models

import (
	"time"
)

type Wallet struct {
	ID        uint    `gorm:"primarykey"`
	UserID    uint    `json:"user_id"`
	Balance   float64 `json:"balance" gorm:"type:decimal(15,2)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"belongsTo"`
}
