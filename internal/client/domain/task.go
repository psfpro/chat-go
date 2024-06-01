package domain

import "github.com/gofrs/uuid"

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
}

func NewTask(id uuid.UUID, title string, description string) *Task {
	return &Task{ID: id, Title: title, Description: description}
}
