package service

import (
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"github.com/FreitasGabriel/chat-app/internal/infra/repository"
)

func NewIUserService(repo repository.IUserRepository) IUserService {
	return &userService{
		repo,
	}
}

type userService struct {
	repo repository.IUserRepository
}

type IUserService interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	ChangePassword(email, oldPassword, newPassword string) error
}
