package handler

import (
	"fmt"

	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"github.com/gin-gonic/gin"
)

func (us *userHandler) CreateUser(c *gin.Context) {
	var user *entity.User

	err := c.BindJSON(&user)
	if err != nil {
		logger.Error("error to bind json", err)
		return
	}

	entityUser := entity.NewUser(user.Name, user.Email, user.Username, user.Password)

	us.service.CreateUser(entityUser)
	c.JSON(200, entityUser)
}

func (us *userHandler) FindByEmail(c *gin.Context) {

	email := c.Query("email")

	if email == "" {
		logger.Info("param email not found")
		c.JSON(201, "param email not found")
		return
	}

	logger.Info(fmt.Sprintf("email: %s", email))

	result, err := us.service.FindByEmail(email)
	if err != nil {
		logger.Error("user not found", err)
		c.JSON(201, "user not found")
		return
	}

	c.JSON(200, result)
}
