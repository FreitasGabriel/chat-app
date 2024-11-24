package handler

import (
	"fmt"

	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/dto"
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"github.com/gin-gonic/gin"
)

func (us *userHandler) CreateUser(c *gin.Context) {
	var user *entity.User

	err := c.BindJSON(&user)
	if err != nil {
		logger.Error("error to bind json", err)
		c.JSON(500, "error to to bind json")
		return
	}

	entityUser := entity.NewUser(user.Name, user.Email, user.Username, user.Password)
	err = us.service.CreateUser(entityUser)
	if err != nil {
		c.JSON(500, "error to create user")
		return
	}

	outputUser := dto.CreateUserOutput{
		Name:     entityUser.Name,
		Email:    entityUser.Email,
		Username: entityUser.Username,
	}

	c.JSON(200, outputUser)
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

func (uh *userHandler) ChangePassword(c *gin.Context) {
	var user dto.ChangePasswordDTO

	err := c.Bind(&user)
	if err != nil {
		logger.Error("error to bind json", err)
		c.JSON(201, "error to bind json")
		return
	}

	err = uh.service.ChangePassword(user.Email, user.OldPassword, user.NewPassword)
	if err != nil {
		logger.Error("error to change password", err)
		c.JSON(201, "error to change password")
		return
	}

	logger.Info("Password changed successfully")
	c.JSON(200, "Password changed successfully")
}
