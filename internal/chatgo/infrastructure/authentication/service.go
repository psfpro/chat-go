package authentication

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const TokenExp = time.Hour * 3

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

type Service struct {
	jwtPrivateKey *rsa.PrivateKey
	jwtPublicKey  *rsa.PublicKey
}

func NewService(jwtPrivateKey *rsa.PrivateKey, jwtPublicKey *rsa.PublicKey) *Service {
	return &Service{jwtPrivateKey: jwtPrivateKey, jwtPublicKey: jwtPublicKey}
}

func (s *Service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (s *Service) CheckPassword(passwordHash string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AccessToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: userID,
	})
	tokenString, err := token.SignedString(s.jwtPrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Service) RefreshToken() (string, error) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *Service) GetUserID(tokenString string) (uuid.UUID, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.jwtPublicKey, nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return claims.UserID, nil
}
