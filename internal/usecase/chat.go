package usecase

import (
	"boilerplate_go_websocket/internal/core"
	"boilerplate_go_websocket/internal/gorm_gen"

	"github.com/labstack/echo/v4"
)

type ChatUseCase interface {
	HandleWebSocket(ctx echo.Context, roomID string, username string) error
}

type chatUseCase struct{
	hubManager *core.HubManager
	query *gorm_gen.Query
}

func NewChatUseCase(hubManager *core.HubManager, query *gorm_gen.Query) ChatUseCase {
	return &chatUseCase{hubManager: hubManager, query: query}
}

// HandleWebSocket implements ChatUseCase.
func (c *chatUseCase) HandleWebSocket(ctx echo.Context, roomID string, username string) error {
	userQuery := c.query.User
	roomQuery := c.query.Room
	user, _ := userQuery.
		WithContext(ctx.Request().Context()).
		Where(userQuery.Username.Eq(username)).
		First()

	room, err := roomQuery.
		WithContext(ctx.Request().Context()).
		Where(roomQuery.ID.Eq(roomID)).
		First()
	if err != nil || room == nil {
		return err
	}

	hub, ok := c.hubManager.GetHub(roomID)
	if !ok && room.HostID == user.ID {
		hub = c.hubManager.InitHub(roomID)
	}

	upgrader := core.DefaultUpgrader()
	core.ServeWs(hub, ctx.Response(), ctx.Request(), upgrader)

	return nil
}
