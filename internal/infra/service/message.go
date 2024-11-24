package service

import (
	"fmt"

	"github.com/FreitasGabriel/chat-app/config/logger"
	"github.com/FreitasGabriel/chat-app/internal/entity"
	"github.com/gorilla/websocket"
)

func ReadMessageFromWebscoket(ws *websocket.Conn, broadcast chan entity.Message) {
	for {
		var message entity.Message
		if err := ws.ReadJSON(&message); err != nil {
			logger.Error("Error to read message", err)
			return
		}

		fmt.Println("message", message)
		broadcast <- message
	}
}

func WriteMessageOnWebsocket(broadcast chan entity.Message, clients map[*websocket.Conn]bool) {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				logger.Error("error to handle message", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
