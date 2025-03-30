package services

import (
	"Shop/internal/models"
	"Shop/internal/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUserProfile(id uint) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}
