package repository

import (
	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/entity"
)

func (r *userRepository) CreateUser(user *entity.User) error {
	r.db.Create(&user)
	return nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("email = ?", email).First(&user)

	if !entity.ValidateID(user.ID) {
		logger.Error("user not found", result.Error)
		return nil, result.Error
	}

	return &user, nil
}
