package repository

import (
	"fmt"

	"github.com/FreitasGabriel/chat-app/internal/entity"
	"gorm.io/gorm"
)

type UserRepositoryInterface struct {
	UserDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryInterface {
	return &UserRepositoryInterface{
		UserDB: db,
	}
}

func (r *UserRepositoryInterface) CreateUser(user entity.User) error {
	fmt.Println("Teste")
	return nil
}
