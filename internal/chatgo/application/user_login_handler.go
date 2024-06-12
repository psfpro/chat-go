package application

import (
	"chatgo/internal/chatgo/domain"
	"context"
	"errors"
)

var ErrUserUnauthorized = errors.New("user unauthorized")

type UserLoginHandler struct {
	userRepository        domain.UserRepository
	authenticationService AuthenticationService
}

func NewUserLoginHandler(userRepository domain.UserRepository, authenticationService AuthenticationService) *UserLoginHandler {
	return &UserLoginHandler{userRepository: userRepository, authenticationService: authenticationService}
}

func (h *UserLoginHandler) Handle(
	ctx context.Context,
	login string,
	password string,
) (accessToken string, refreshToken string, err error) {
	user, err := h.userRepository.GetByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			err = ErrUserUnauthorized
			return
		}
		return
	}
	if err = h.authenticationService.CheckPassword(user.PasswordHash, password); err != nil {
		return
	}

	accessToken, err = h.authenticationService.AccessToken(user.ID)
	if err != nil {
		return
	}
	refreshToken, err = h.authenticationService.RefreshToken()
	if err != nil {
		return
	}

	return accessToken, refreshToken, nil
}
