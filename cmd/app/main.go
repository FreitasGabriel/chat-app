package main

import (
	"time"

	configs "github.com/FreitasGabriel/chat-app/config"
	database "github.com/FreitasGabriel/chat-app/config/database/postgres"
	config "github.com/FreitasGabriel/chat-app/config/dependencies"
	"github.com/FreitasGabriel/chat-app/config/logger"
	redisConfig "github.com/FreitasGabriel/chat-app/config/redis"
	"github.com/FreitasGabriel/chat-app/internal/infra/handler"
	"github.com/FreitasGabriel/chat-app/internal/infra/service"
	"github.com/FreitasGabriel/chat-app/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	logger.Info("Starting server...")

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	rateLimit := redisConfig.NewRateLimiter(redisClient, 1*time.Minute, 5)

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
	rlHandler := handler.RateLimitMiddleware(rateLimit, nil)

	gin.SetMode(gin.ReleaseMode)
	c := gin.Default()

	c.Use(rlHandler)
	c.Use(func(c *gin.Context) {
		gin.Recovery()
		c.Set("jwt_secret", conf.JWTSecret)
		c.Set("jwt_expires_in", conf.JWTExpiresIn)
		c.Set("cypher_key", conf.CypherKey)
		c.Next()
	})

	routes.InitRoutes(&c.RouterGroup, websocketBroadcast, userHandler)

	go service.WriteMessageOnWebsocket(websocketBroadcast.Broadcast, websocketBroadcast.Clients, []byte(conf.CypherKey))

	if err := c.Run(":8000"); err != nil {
		logger.Error("Error to start server", err)
	}
}
