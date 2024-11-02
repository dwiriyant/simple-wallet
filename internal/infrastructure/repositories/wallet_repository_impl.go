package repositories

import (
	"simple-wallet/internal/domain/models"
	"simple-wallet/internal/domain/repositories"

	"gorm.io/gorm"
)

type walletRepositoryImpl struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) repositories.WalletRepository {
	return &walletRepositoryImpl{db: db}
}

func (r *walletRepositoryImpl) CreateWallet(wallet *models.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *walletRepositoryImpl) GetByUserID(userID uint) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepositoryImpl) UpdateBalance(wallet *models.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *walletRepositoryImpl) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
