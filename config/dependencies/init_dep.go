package config

import (
	"github.com/FreitasGabriel/chat-app/internal/infra/handler"
	"github.com/FreitasGabriel/chat-app/internal/infra/repository"
	"gorm.io/gorm"
)

func InitDependencies(gormDB *gorm.DB) handler.IUserHandler {
	userDB := repository.NewUserRepository(gormDB)
	userHandler := handler.NewUserHandler(userDB)
	return userHandler
}
