package main

import (
	configs "github.com/FreitasGabriel/chat-app/config"
	database "github.com/FreitasGabriel/chat-app/config/database/postgres"
	config "github.com/FreitasGabriel/chat-app/config/dependencies"
	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/infra/webserver/handler"
	"github.com/FreitasGabriel/chat-app/internal/routes"
	"github.com/FreitasGabriel/chat-app/internal/service"
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
	routes.InitRoutes(&c.RouterGroup, websocketBroadcast, userHandler)

	go service.WriteMessageOnWebsocket(websocketBroadcast.Broadcast, websocketBroadcast.Clients)

	if err := c.Run(":8000"); err != nil {
		logger.Error("Error to start server", err)
	}
}
