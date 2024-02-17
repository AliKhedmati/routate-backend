package services

import (
	"context"
	"github.com/AliKhedmati/routate-backend/src/models"
	"github.com/AliKhedmati/routate-backend/src/repositories"
)

// UserService represents a services for user-related operations.
type UserService struct {
	userRepository *repositories.UserRepository
}

// NewUserService Returns an instance of UserService
func NewUserService(UserRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: UserRepository,
	}
}

// Create Creates a new user.
func (service *UserService) Create(ctx context.Context, user *models.User) error {
	return service.userRepository.Create(ctx, user)
}

// FindByID Finds a user by their ID.
func (service *UserService) FindByID(ctx context.Context, id string) (*models.User, error) {
	return service.userRepository.FindByID(ctx, id)
}
