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
