package entity

import (
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMessage(username, message string) *Message {
	return &Message{
		Username: username,
		Message:  message,
	}
}
