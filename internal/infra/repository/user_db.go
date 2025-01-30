package repository

import (
	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/dto"
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

func (r *userRepository) FindByEmail(email string) (*entity.User, *dto.FindUserOutput, error) {
	var user entity.User
	result := r.db.Where("email = ?", email).First(&user)

	if !entity.ValidateID(user.ID) {
		logger.Error("user not found", result.Error)
		return nil, nil, result.Error
	}

	var userDTO = dto.FindUserOutput{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}

	return &user, &userDTO, nil
}

func (r *userRepository) ChangePassword(email, oldPassword, newPassword string) error {
	user, _, err := r.FindByEmail(email)
	if err != nil {
		logger.Error("user not found", err)
		return err
	}

	isValid, err := user.ValidatePassword(oldPassword)
	if !isValid {
		logger.Error("invalid old password", err)
		return err
	}

	hashedPassword, _ := entity.GenerateHashedPassword(newPassword)

	result := r.db.Model(&user).Update("password", hashedPassword)
	if result.Error != nil {
		logger.Error("error to change password", result.Error)
		return result.Error
	}

	return nil
}
