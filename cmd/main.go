package main

import (
	"net/http"

	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/handler"
	"github.com/FreitasGabriel/chat-app/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Starting server...")

	websocketBroadcast := handler.NewWebsocketBroadcast()

	gin.SetMode(gin.ReleaseMode)
	c := gin.Default()

	c.GET("/health", handler.Health)
	c.GET("/ws", websocketBroadcast.HandleWebsocketConnection)
	c.StaticFS("/chat", http.Dir("./templates"))

	go service.WriteMessageOnWebsocket(websocketBroadcast.Broadcast, websocketBroadcast.Clients)

	if err := c.Run(":8000"); err != nil {
		logger.Error("Error to start server", err)
	}
}
