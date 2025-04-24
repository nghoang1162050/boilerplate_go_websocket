package usecase

import (
	// "boilerplate_go/internal/model"
	// "boilerplate_go/internal/repository"
	"boilerplate_go_websocket/internal/dto"
	"context"
)

type AuthUseCase interface {
	Login(ctx context.Context, username, password string) (dto.BaseResponse, error)
}

type authUseCase struct {
	// repo repository.BaseRepository[model.User]
	// userRepo repository.UserRepository
}

func NewAuthUseCase() AuthUseCase {
	return &authUseCase{}
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(ctx context.Context, username string, password string) (dto.BaseResponse, error) {
	// existingUser, err := a.repo.First("username = ?", username)
	// if err != nil || existingUser == nil {
	// 	return dto.NewBaseResponse(400, "Invalid user info.", nil), err
	// }

	// if !utils.CheckPasswordHash(password, existingUser.PasswordHash) {
	// 	return dto.NewBaseResponse(400, "Invalid user info.", nil), nil
	// }

	// // TODO: hardcoded role, should be replaced with a proper role management system.
	// tokenString, expired, err := utils.GenerateJWTToken(username, []string{"admin"})
	// if err != nil {
	// 	return dto.NewBaseResponse(500, "Failed to generate token", nil), err
	// }

	// loginResponse := dto.UserLoginResponse{
	// 	Username: username,
	// 	Role:     "admin",
	// 	Token:    tokenString,
	// 	Expiry:   expired,
	// }

	// return dto.NewBaseResponse(200, "Login successful", loginResponse), nil

	panic("Login method not implemented")
}
