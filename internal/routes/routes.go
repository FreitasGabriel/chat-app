package routes

import (
	"github.com/FreitasGabriel/chat-app/internal/infra/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(c *gin.RouterGroup, broadCastHandler *handler.WebsocketBroadcast, userHandler handler.IUserHandler) {
	c.GET("/health", handler.Health)
	c.GET("/ws", broadCastHandler.HandleWebsocketConnection)
	c.POST("/user", userHandler.CreateUser)
	c.PUT("/user/password", userHandler.ChangePassword)
	c.GET("/user", userHandler.FindByEmail)
}
