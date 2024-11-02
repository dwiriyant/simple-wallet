package usecases

import (
	"errors"
	"simple-wallet/internal/domain/models"
	"simple-wallet/internal/domain/repositories"
)

type WalletService struct {
	walletRepo repositories.WalletRepository
	userRepo   repositories.UserRepository
}

func NewWalletService(walletRepo repositories.WalletRepository, userRepo repositories.UserRepository) *WalletService {
	return &WalletService{walletRepo: walletRepo, userRepo: userRepo}
}

func (s *WalletService) Create(userID uint) error {
	wallet := &models.Wallet{
		UserID:  userID,
		Balance: 50000,
	}
	return s.walletRepo.CreateWallet(wallet)
}

func (s *WalletService) GetBalance(userID uint) (float64, error) {
	wallet, err := s.walletRepo.GetByUserID(userID)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func (s *WalletService) Transfer(senderUserID uint, recipientUsername string, amount float64) error {
	tx := s.walletRepo.BeginTransaction()
	defer func() {
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
		}
	}()

	senderWallet, err := s.walletRepo.GetByUserID(senderUserID)
	if err != nil {
		return err
	}

	recipient, err := s.userRepo.GetByUsername(recipientUsername)
	if err != nil {
		return err
	}

	recipientWallet, err := s.walletRepo.GetByUserID(recipient.ID)
	if err != nil {
		return err
	}

	if senderWallet.Balance < amount {
		return errors.New("insufficient balance")
	}

	senderWallet.Balance -= amount
	recipientWallet.Balance += amount

	if err := s.walletRepo.UpdateBalance(senderWallet); err != nil {
		return err
	}
	if err := s.walletRepo.UpdateBalance(recipientWallet); err != nil {
		return err
	}

	tx.Commit()

	return nil
}
