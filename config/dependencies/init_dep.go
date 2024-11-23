package config

import (
	"github.com/FreitasGabriel/chat-app/internal/infra/repository"
	"github.com/FreitasGabriel/chat-app/internal/infra/webserver/handler"
	"gorm.io/gorm"
)

func InitDependencies(gormDB *gorm.DB) *handler.UserHandlerinterface {
	userDB := repository.NewUserRepository(gormDB)
	userHandler := handler.NewUserHandler(userDB)
	return userHandler
}
