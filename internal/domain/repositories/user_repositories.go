package repositories

import "simple-wallet/internal/domain/models"

type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
	Create(user *models.User) error
	Login(username string, password string) error
	HashPassword(password string, user *models.User) error
	CheckPassword(password string, user *models.User) error
}
