package controller

import (
	"boilerplate_go_websocket/internal/dto"
	"boilerplate_go_websocket/internal/usecase"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(ctx echo.Context) error
}

type authController struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) AuthController {
	return &authController{authUseCase: authUseCase}
}

// Login implements AuthController.
// @Summary Login user
// @Description Authenticate a user and return a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.UserLoginDTO true "User login data"
// @Success 200 {object} dto.BaseResponse
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /auth/login [post]
func (a *authController) Login(ctx echo.Context) error {
	var userDto dto.UserLoginDTO
	if err := ctx.Bind(&userDto); err != nil {
		return ctx.JSON(400, dto.NewBaseResponse(400, "Invalid request", nil))
	}

	response, _ := a.authUseCase.Login(ctx.Request().Context(), userDto.Username, userDto.Password)

	return ctx.JSON(response.Code, response)
}
