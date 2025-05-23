package entity

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	WriteBufferPool: nil,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
