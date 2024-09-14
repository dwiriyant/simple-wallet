package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Wallet struct {
	ID        uint    `gorm:"primarykey"`
	UserID    uint    `json:"user_id"`
	Balance   float64 `json:"balance" gorm:"type:decimal(15,2)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"belongsTo"`
}

func (w *Wallet) Transfer(db *gorm.DB, amount float64) (err error) {
	tx := db.Begin()

	defer func() {
		switch err {
		case nil:
			err = tx.Commit().Error
		default:
			tx.Rollback()
		}
	}()

	if err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&w, w.ID).Error; err != nil {
		return
	}

	if w.Balance < amount {
		return errors.New("Insufficient funds")
	}

	w.Balance -= amount
	if err = tx.Save(&w).Error; err != nil {
		return
	}

	return
}
