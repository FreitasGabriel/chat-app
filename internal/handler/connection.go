package handler

import (
	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"github.com/FreitasGabriel/chat-app/internal/service"
	"github.com/gin-gonic/gin"
)

func (wb *WebsocketBroadcast) HandleWebsocketConnection(c *gin.Context) {
	ws, err := entity.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("Error to upgrade connection", err)
		return
	}
	defer ws.Close()

	wb.Clients[ws] = true

	service.ReadMessageFromWebscoket(ws, wb.Broadcast)
}
