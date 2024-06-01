package application

import "github.com/gofrs/uuid"

//go:generate mockery --name ChatGoService --with-expecter
type ChatGoService interface {
	AddTask(title string, description string) (uuid.UUID, error)
}
