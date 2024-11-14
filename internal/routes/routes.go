package routes

import (
	"net/http"

	"github.com/FreitasGabriel/chat-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func Init(c *gin.RouterGroup, handlerBroadcast *handler.WebsocketBroadcast) {
	c.GET("/health", handler.Health)
	c.GET("ws", handlerBroadcast.HandleWebsocketConnection)
	c.StaticFS("/chat/", http.Dir("../../templates/index.html"))
}
