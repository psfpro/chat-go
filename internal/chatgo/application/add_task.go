package application

import (
	"chatgo/internal/chatgo/domain"
	"context"
	"github.com/gofrs/uuid"
)

type AddTask struct {
	worker         Worker
	taskRepository domain.TaskRepository
}

func NewAddTask(worker Worker, taskRepository domain.TaskRepository) *AddTask {
	return &AddTask{worker: worker, taskRepository: taskRepository}
}

func (h AddTask) Handle(ctx context.Context, userID uuid.UUID, title string, description string) (uuid.UUID, error) {
	id := uuid.Must(uuid.NewV6())
	task := domain.NewTask(id, userID, title, description)
	h.worker.Do(task)
	err := h.taskRepository.Save(ctx, task)
	if err != nil {
		return id, err
	}

	return id, nil
}
