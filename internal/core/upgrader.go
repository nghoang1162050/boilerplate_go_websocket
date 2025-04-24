package core

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func DefaultUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// Optionally add a CheckOrigin function.
		CheckOrigin: func(r *http.Request) bool { return true },
	}
}
