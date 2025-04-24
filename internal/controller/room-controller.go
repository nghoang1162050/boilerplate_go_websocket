package controller

import (
	"boilerplate_go_websocket/internal/dto"
	"boilerplate_go_websocket/internal/usecase"
	"boilerplate_go_websocket/internal/utils"

	"github.com/labstack/echo/v4"
)

type RoomController interface {
	CreateRoom(ctx echo.Context) error
	GetRoom(ctx echo.Context) error
	CloseRoom(ctx echo.Context) error
}

type roomController struct {
	roomUseCase usecase.RoomUseCase
}

func NewRoomController(roomUseCase usecase.RoomUseCase) RoomController {
	return &roomController{roomUseCase: roomUseCase}
}

// CreateRoom implements RoomController.
func (r *roomController) CreateRoom(ctx echo.Context) error {
	roomDTO := new(dto.RoomDTO)
	if err := ctx.Bind(roomDTO); err != nil {
		return ctx.JSON(400, dto.NewBaseResponse(400, "invalid request", nil))
	}

	// get current user from token
	username, _ := utils.ExtractUsernameFromToken(ctx.Request().Header.Get("Authorization"))

	roomID, err := r.roomUseCase.CreateRoom(ctx, roomDTO, username)
	if err != nil {
		return ctx.JSON(500, dto.NewBaseResponse(500, "failed to create room", nil))
	}

	response := dto.NewBaseResponse(200, "room created", roomID)
	return ctx.JSON(response.Code, response)
}

// CloseRoom implements RoomController.
func (r *roomController) CloseRoom(ctx echo.Context) error {
	roomID := ctx.Param("room_id")
	response, _ := r.roomUseCase.CloseRoom(ctx, roomID)
	
	return ctx.JSON(response.Code, response)
}

// GetRoom implements RoomController.
func (r *roomController) GetRoom(ctx echo.Context) error {
	roomID := ctx.Param("room_id")
	response, _ := r.roomUseCase.GetRoom(ctx, roomID)

	return ctx.JSON(response.Code, response)
}