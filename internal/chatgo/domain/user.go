package domain

import "github.com/gofrs/uuid"

type User struct {
	ID           uuid.UUID
	Login        string
	PasswordHash string
}

func NewUser(ID uuid.UUID, login string, passwordHash string) *User {
	return &User{ID: ID, Login: login, PasswordHash: passwordHash}
}
