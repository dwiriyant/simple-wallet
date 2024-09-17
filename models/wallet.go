package models

import (
	"errors"
	"simple-wallet/db"
	"time"

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

func (w *Wallet) Transfer(toWallet Wallet, amount float64) (err error) {
	tx := db.DB.Begin()

	defer func() {
		switch err {
		case nil:
			err = tx.Commit().Error
		default:
			tx.Rollback()
		}
	}()

	if err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&w).Error; err != nil {
		return
	}

	if w.Balance < amount {
		return errors.New("Insufficient funds")
	}

	w.Balance -= amount
	if err = tx.Model(&w).Update("balance", w.Balance).Error; err != nil {
		return
	}

	toWallet.Balance += amount
	if err = tx.Model(&toWallet).Update("balance", toWallet.Balance).Error; err != nil {
		return errors.New("Failed to update recipient wallet")
	}

	return
}
