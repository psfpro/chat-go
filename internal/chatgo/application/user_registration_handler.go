package application

import (
	"chatgo/internal/chatgo/domain"
	"context"
	"errors"
	"github.com/gofrs/uuid"
)

var ErrUserLoginAlreadyTaken = errors.New("user login already taken")

type UserRegistrationHandler struct {
	userRepository        domain.UserRepository
	authenticationService AuthenticationService
}

func NewUserRegistrationHandler(userRepository domain.UserRepository, authenticationService AuthenticationService) *UserRegistrationHandler {
	return &UserRegistrationHandler{userRepository: userRepository, authenticationService: authenticationService}
}

func (h *UserRegistrationHandler) Handle(
	ctx context.Context,
	login string,
	password string,
) (accessToken string, refreshToken string, err error) {
	userID := uuid.Must(uuid.NewV6())
	user, err := h.userRepository.GetByLogin(ctx, login)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return
	}
	if user != nil {
		err = ErrUserLoginAlreadyTaken
		return
	}
	passwordHash, err := h.authenticationService.HashPassword(password)
	if err != nil {
		return
	}

	user = domain.NewUser(
		userID,
		login,
		passwordHash,
	)
	if err = h.userRepository.Save(ctx, user); err != nil {
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
