package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
