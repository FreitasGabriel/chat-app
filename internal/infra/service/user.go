package service

import (
	"github.com/FreitasGabriel/chat-app/internal/entity"
)

func (us *userService) CreateUser(user *entity.User) error {
	err := us.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) FindByEmail(email string) (*entity.User, error) {
	user, err := us.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) ChangePassword(email, oldPassword, newPassword string) error {
	err := us.repo.ChangePassword(email, oldPassword, newPassword)
	if err != nil {
		return err
	}
	return nil
}
