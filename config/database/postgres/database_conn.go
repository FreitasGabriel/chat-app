package database

import (
	"fmt"

	configs "github.com/FreitasGabriel/chat-app/config"
	"github.com/FreitasGabriel/chat-app/config/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(config *configs.Conf) (*gorm.DB, error) {

	host := config.DBHost
	user := config.DBUser
	password := config.DBPassword
	port := config.DBPort
	dbName := config.DBName

	dsn := NewDSN(host, user, password, port, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Error to connect to database", err)
		return nil, err
	}

	logger.Info("Database connected")

	return db, nil
}

func NewDSN(host, user, password, port, dbName string) string {
	newDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	return newDSN
}
