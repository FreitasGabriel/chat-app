package main

import (
	configs "github.com/FreitasGabriel/chat-app/config"
	database "github.com/FreitasGabriel/chat-app/config/database/postgres"
	config "github.com/FreitasGabriel/chat-app/config/dependencies"
	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/infra/handler"
	"github.com/FreitasGabriel/chat-app/internal/infra/service"
	"github.com/FreitasGabriel/chat-app/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Starting server...")

	websocketBroadcast := handler.NewWebsocketBroadcast()

	conf, err := configs.LoadConfig(".")
	if err != nil {
		logger.Error("Error to load config", err)
		panic(err)
	}

	gormDB, err := database.NewDatabaseConnection(conf)
	if err != nil {
		logger.Error("Error to connect to database", err)
		panic(err)
	}

	userHandler := config.InitDependencies(gormDB)

	gin.SetMode(gin.ReleaseMode)
	c := gin.Default()
	c.Use(gin.Recovery())

	c.Use(func(c *gin.Context) {
		c.Set("jwt_secret", conf.JWTSecret)
		c.Set("jwt_expires_in", conf.JWTExpiresIn)
		c.Next()
	})

	routes.InitRoutes(&c.RouterGroup, websocketBroadcast, userHandler)

	go service.WriteMessageOnWebsocket(websocketBroadcast.Broadcast, websocketBroadcast.Clients)

	if err := c.Run(":8000"); err != nil {
		logger.Error("Error to start server", err)
	}
}
