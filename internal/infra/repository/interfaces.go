package repository

import (
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db,
	}
}

type userRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
