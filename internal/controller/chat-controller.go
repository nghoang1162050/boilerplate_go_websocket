package controller

import (
	"boilerplate_go_websocket/internal/usecase"
	"boilerplate_go_websocket/internal/utils"

	"github.com/labstack/echo/v4"
)

type ChatController interface {
	HandleWebSocket(ctx echo.Context) error
}

type chatController struct {
	chatUseCase usecase.ChatUseCase
}

func NewChatController(chatUseCase usecase.ChatUseCase) ChatController {
	return &chatController{chatUseCase: chatUseCase}
}

// HandleWebSocket implements ChatController.
func (c *chatController) HandleWebSocket(ctx echo.Context) error {
	hubID := ctx.Param("hubID")
	if hubID == "" {
		return echo.NewHTTPError(400, "hubID is required")
	}

	username, _ := utils.ExtractUsernameFromToken(ctx.Request().Header.Get("Authorization"))

	err := c.chatUseCase.HandleWebSocket(ctx, hubID, username)

	return err
}
