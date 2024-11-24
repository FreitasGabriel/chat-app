package entity

import (
	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/google/uuid"
)

func NewUUID() uuid.UUID {
	return uuid.New()
}

func ValidateID(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("Error to parse UUID", err)
		return false
	}
	return true
}
