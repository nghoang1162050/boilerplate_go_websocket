package router

import (
	"boilerplate_go_websocket/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewRoomRouter(e *echo.Group, r controller.RoomController) {
	roomGroup := e.Group("/room")

	roomGroup.POST("", r.CreateRoom)
	roomGroup.GET("/:room_id", r.GetRoom)
	roomGroup.PATCH("/:room_id/close", r.CloseRoom)
}
