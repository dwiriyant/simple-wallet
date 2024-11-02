package repositories

import (
	"simple-wallet/internal/domain/models"

	"gorm.io/gorm"
)

type WalletRepository interface {
	CreateWallet(wallet *models.Wallet) error
	GetByUserID(userID uint) (*models.Wallet, error)
	UpdateBalance(wallet *models.Wallet) error
	BeginTransaction() *gorm.DB
}
