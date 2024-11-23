package routes

import (
	"github.com/FreitasGabriel/chat-app/internal/infra/webserver/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(c *gin.RouterGroup, broadCastHandler *handler.WebsocketBroadcast, userHandler *handler.UserHandlerinterface) {
	c.GET("/health", handler.Health)
	c.GET("/ws", broadCastHandler.HandleWebsocketConnection)
	c.POST("/user", userHandler.CreateUserHandler)
}
