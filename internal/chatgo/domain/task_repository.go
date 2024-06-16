package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

//go:generate mockery --name TaskRepository --with-expecter
type TaskRepository interface {
	GetAllByUserId(ctx context.Context, userID uuid.UUID) ([]*Task, error)
	Save(ctx context.Context, task *Task) error
}
