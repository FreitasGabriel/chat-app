package handler

import (
	"github.com/FreitasGabriel/chat-app/internal/infra/repository"
	"github.com/gin-gonic/gin"
)

type UserHandlerinterface struct {
	UserDB *repository.UserRepositoryInterface
}

func NewUserHandler(db *repository.UserRepositoryInterface) *UserHandlerinterface {
	return &UserHandlerinterface{
		UserDB: db,
	}
}

func (us *UserHandlerinterface) CreateUserHandler(c *gin.Context) {

}
