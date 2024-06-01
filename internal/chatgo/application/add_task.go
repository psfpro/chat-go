package application

import (
	"chatgo/internal/chatgo/domain"
	"github.com/gofrs/uuid"
)

type AddTask struct {
	worker Worker
}

func NewAddTask(worker Worker) *AddTask {
	return &AddTask{worker: worker}
}

func (h AddTask) Handle(title string, description string) (uuid.UUID, error) {
	id := uuid.Must(uuid.NewV6())
	task := domain.NewTask(id, title, description)
	h.worker.Do(task)

	return id, nil
}
