package dumb

import "github.com/gofrs/uuid"

type ChatGoService struct {
}

func (s ChatGoService) AddTask(title string, description string) (uuid.UUID, error) {
	return uuid.Must(uuid.NewV6()), nil
}
