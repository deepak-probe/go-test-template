package service

import (
	"context"
	model "user-api/internal/user"
	"user-api/internal/user/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepository: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	// You can add business logic here, like hashing the password
	return s.userRepository.CreateUser(ctx, user)
}
