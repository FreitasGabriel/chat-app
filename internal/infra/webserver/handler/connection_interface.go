package handler

import (
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"github.com/gorilla/websocket"
)

type WebsocketBroadcast struct {
	Broadcast chan entity.Message
	Clients   map[*websocket.Conn]bool
}

func NewWebsocketBroadcast() *WebsocketBroadcast {
	return &WebsocketBroadcast{
		Broadcast: make(chan entity.Message),
		Clients:   make(map[*websocket.Conn]bool),
	}
}
