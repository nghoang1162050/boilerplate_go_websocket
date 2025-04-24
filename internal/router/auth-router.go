package router

import (
	"boilerplate_go_websocket/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewAuthRouter(e *echo.Group, a controller.AuthController) {
	authGroup := e.Group("/auth")

	authGroup.POST("/login", a.Login)
}
