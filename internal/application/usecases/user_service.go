package usecases

import (
	"simple-wallet/internal/domain/models"
	"simple-wallet/internal/domain/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetByUsername(username string) (*models.User, error) {
	return s.userRepo.GetByUsername(username)
}

func (s *UserService) Register(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) Login(username string, password string) error {
	return s.userRepo.Login(username, password)
}
