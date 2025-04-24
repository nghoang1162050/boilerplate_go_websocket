package core

import (
	"boilerplate_go_websocket/internal/constants"
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub *Hub // The Hub to which the client belongs

	// The WebSocket Connection for the client
	Connection *websocket.Conn

	// Channel for sending messages to the client
	Send chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Connection.Close()
	}()

	c.Connection.SetReadLimit(constants.MaxMessageSize)
	c.Connection.SetReadDeadline(time.Now().Add(constants.PongWait))
	c.Connection.SetPongHandler(func(string) error {
		c.Connection.SetReadDeadline(time.Now().Add(constants.PongWait))
		return nil
	})

	msgChan := make(chan []byte)
	errChan := make(chan error, 1)

	go func() {
		defer close(msgChan)
		for {
			_, message, err := c.Connection.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			message = bytes.TrimSpace(bytes.Replace(message, constants.Newline, constants.Space, -1))
            msgChan <- message
		}
	}()

	for {
		select {
		case msg, ok := <-msgChan:
			if !ok {
				return
			}
			c.Hub.Broadcast <- msg

		case err := <-errChan:
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(constants.PingPeriod)

	defer func() {
		ticker.Stop()
		c.Connection.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Connection.SetWriteDeadline(time.Now().Add(constants.WriteWait))
			if !ok {
				c.Connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Connection.SetWriteDeadline(time.Now().Add(constants.WriteWait))
			if err := c.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
