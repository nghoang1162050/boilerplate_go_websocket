package router

import (
	"boilerplate_go_websocket/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewChatRouter(e *echo.Group, r controller.ChatController) {
	chatGroup := e.Group("/ws")

	chatGroup.GET("/:hubID", r.HandleWebSocket)
}
