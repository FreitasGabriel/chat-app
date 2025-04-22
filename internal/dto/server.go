package dto

import "github.com/FreitasGabriel/chat-app/internal/entity"

type Server struct {
	ID       string
	Name     string
	Channels []Channel
	Users    []entity.User
}
