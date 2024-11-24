package handler

import (
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"github.com/FreitasGabriel/chat-app/internal/infra/service"
	"github.com/gin-gonic/gin"
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

func NewUserHandler(service service.IUserService) IUserHandler {
	return &userHandler{
		service,
	}
}

type userHandler struct {
	service service.IUserService
}

type IUserHandler interface {
	CreateUser(c *gin.Context)
	FindByEmail(c *gin.Context)
	ChangePassword(c *gin.Context)
	UserLogin(c *gin.Context)
}
