package config

import (
	"github.com/FreitasGabriel/chat-app/internal/infra/handler"
	"github.com/FreitasGabriel/chat-app/internal/infra/repository"
	"github.com/FreitasGabriel/chat-app/internal/infra/service"
	"gorm.io/gorm"
)

func InitDependencies(gormDB *gorm.DB) handler.IUserHandler {
	userDB := repository.NewUserRepository(gormDB)
	userService := service.NewIUserService(userDB)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}
