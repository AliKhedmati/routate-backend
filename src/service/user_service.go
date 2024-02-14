package service

import (
	"context"
	"github.com/AliKhedmati/routate-backend/src/model"
	"github.com/AliKhedmati/routate-backend/src/repository"
)

// UserService represents a service for user-related operations.
type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService initializes a new instance of UserService.
func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

// Create creates a new user.
func (service *UserService) Create(ctx context.Context, user *model.User) error {
	return service.userRepository.Create(ctx, user)
}

// FindByID finds a user by their ID.
func (service *UserService) FindByID(ctx context.Context, id string) (*model.User, error) {
	return service.userRepository.FindByID(ctx, id)
}
