package repository

import (
	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/entity"
)

func (r *userRepository) CreateUser(user *entity.User) error {
	err := r.db.Create(&user)
	if err.Error != nil {
		logger.Error("error to create user", err.Error)
		return err.Error
	}
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

func (r *userRepository) ChangePassword(email, oldPassword, newPassword string) error {
	user, err := r.FindByEmail(email)
	if err != nil {
		logger.Error("user not found", err)
		return err
	}

	isValid, err := user.ValidatePassword(oldPassword)
	if !isValid {
		logger.Error("invalid old password", err)
		return err
	}

	result := r.db.Model(&user).Update("password", newPassword)
	if result.Error != nil {
		logger.Error("error to change password", result.Error)
		return result.Error
	}

	return nil
}
