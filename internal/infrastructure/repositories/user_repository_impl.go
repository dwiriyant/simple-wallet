package repositories

import (
	"simple-wallet/internal/domain/models"
	"simple-wallet/internal/domain/repositories"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) Login(username string, password string) error {
	var user models.User
	if err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return err
	}
	return nil
}
