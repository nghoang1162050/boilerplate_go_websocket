package core

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Hub struct {
	// Registered clients
	Clients map[*Client]bool
	// Broadcast channel for messages to be sent to clients
	Broadcast chan []byte
	// Register channel for new clients to be registered
	Register chan *Client
	// Close channel for closing connections
	CloseChan chan struct{}
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		CloseChan: make(chan struct{}),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		case <-h.CloseChan:
			log.Println("Hub is closing")
			for client := range h.Clients {
				close(client.Send)
				delete(h.Clients, client)
			}
			return
		}
	}
}

func (h *Hub) Close() {
    close(h.CloseChan)
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	client := &Client{
		Hub:        hub,
		Connection: conn,
		Send:       make(chan []byte, 256),
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
