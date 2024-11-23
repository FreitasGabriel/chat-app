package repository

import "github.com/FreitasGabriel/chat-app/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
