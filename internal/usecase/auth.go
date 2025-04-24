package usecase

import (
	"boilerplate_go_websocket/internal/dto"
	"boilerplate_go_websocket/internal/gorm_gen"
	"boilerplate_go_websocket/internal/utils"
	"context"
)

type AuthUseCase interface {
	Login(ctx context.Context, username, password string) (dto.BaseResponse, error)
}

type authUseCase struct {
	// repo repository.BaseRepository[model.User]
	// userRepo repository.UserRepository
	query *gorm_gen.Query
}

func NewAuthUseCase(query *gorm_gen.Query) AuthUseCase {
	return &authUseCase{query: query}
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(ctx context.Context, username string, password string) (dto.BaseResponse, error) {
	u := a.query.User
	existingUser, err := u.WithContext(ctx).Where(u.Username.Eq(username)).First()
	if existingUser == nil || err != nil {
		return dto.NewBaseResponse(400, "Invalid user info.", nil), nil
	}

	if !utils.CheckPasswordHash(password, existingUser.PasswordHash) {
		return dto.NewBaseResponse(400, "Invalid user info.", nil), nil
	}

	// TODO: hardcoded role, should be replaced with a proper role management system.
	tokenString, expired, err := utils.GenerateJWTToken(username, []string{"admin"})
	if err != nil {
		return dto.NewBaseResponse(500, "Failed to generate token", nil), err
	}

	loginResponse := dto.UserLoginResponse{
		Username: username,
		Role:     "admin",
		Token:    tokenString,
		Expiry:   expired,
	}

	return dto.NewBaseResponse(200, "Login successful", loginResponse), nil
}
