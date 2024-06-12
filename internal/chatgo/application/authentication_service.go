package application

import (
	"github.com/gofrs/uuid"
)

//go:generate mockery --name AuthenticationService --with-expecter
type AuthenticationService interface {
	HashPassword(password string) (string, error)
	CheckPassword(passwordHash string, providedPassword string) error
	AccessToken(userID uuid.UUID) (string, error)
	RefreshToken() (string, error)
	GetUserID(tokenString string) (uuid.UUID, error)
}
