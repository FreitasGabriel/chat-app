package entity

import "github.com/gorilla/websocket"

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
