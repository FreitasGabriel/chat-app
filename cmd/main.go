package main

import (
	"fmt"

	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	logger.Info("Starting server...")

	c := gin.Default()

	routes.Init(&c.RouterGroup)

	if err := c.Run(":8000"); err != nil {
		logger.Error("Error to start server", err)
	}
}
