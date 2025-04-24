package usecase

import (
	"boilerplate_go_websocket/internal/dto"
	"boilerplate_go_websocket/internal/gorm_gen"
	"boilerplate_go_websocket/internal/model"
	"boilerplate_go_websocket/internal/utils"
	"errors"

	"github.com/labstack/echo/v4"
)

type RoomUseCase interface {
	CreateRoom(ctx echo.Context, roomDTO *dto.RoomDTO, username string) (string, error)
	GetRoom(ctx echo.Context, roomID string) (dto.BaseResponse, error)
	CloseRoom(ctx echo.Context, roomID string) (dto.BaseResponse, error)
}

type roomUseCase struct {
	query *gorm_gen.Query
}

func NewRoomUseCase(query *gorm_gen.Query) RoomUseCase {
	return &roomUseCase{
		query: query,
	}
}

// CreateRoom implements RoomUseCase.
func (r *roomUseCase) CreateRoom(ctx echo.Context, roomDTO *dto.RoomDTO, username string) (string, error) {
	// get userId from username
	user, err := r.query.User.WithContext(ctx.Request().Context()).Where(r.query.User.Username.Eq(username)).First()
	if err != nil {
		return "not fould username", err
	}

	roomID, _ := utils.GenerateRoomID(user.Username)

	if err := r.query.Room.WithContext(ctx.Request().Context()).Create(&model.Room{
		ID:          roomID,
		Name:        roomDTO.Name,
		Description: roomDTO.Desc,
		HostID:      user.ID,
	}); err != nil {
		return "", err
	}

	return roomID, nil
}

// GetRoom implements RoomUseCase.
func (r *roomUseCase) GetRoom(ctx echo.Context, roomID string) (dto.BaseResponse, error) {
	currentRoom, err := r.query.Room.WithContext(ctx.Request().Context()).Where(r.query.Room.ID.Eq(roomID)).First()
	if err != nil {
		return dto.NewBaseResponse(404, "room not found", nil), err
	}

	return dto.NewBaseResponse(200, "room found", currentRoom), nil
}

// CloseRoom implements RoomUseCase.
func (r *roomUseCase) CloseRoom(ctx echo.Context, roomID string) (dto.BaseResponse, error) {
	currentRoom, _ := r.query.Room.WithContext(ctx.Request().Context()).Where(r.query.Room.ID.Eq(roomID)).First()
    if currentRoom == nil {
        return dto.NewBaseResponse(400, "not fould room", nil), errors.New("room not found")
    }

	if len(currentRoom.IsClosed) > 0 && currentRoom.IsClosed[0] == 1 {
        return dto.NewBaseResponse(400, "room is already closed", nil), errors.New("room already closed")
    }

	room := r.query.Room
	if _, err := room.WithContext(ctx.Request().Context()).Where(room.ID.Eq(roomID)).UpdateColumn(room.IsClosed, 1); err != nil {
		return dto.NewBaseResponse(500, "internal server error", err.Error()), err
	}

	return dto.NewBaseResponse(200, "", nil), nil
}
