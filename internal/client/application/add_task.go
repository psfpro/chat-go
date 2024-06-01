package application

import (
	"chatgo/internal/client/domain"
	"github.com/gofrs/uuid"
)

type AddTask struct {
	repository domain.TaskRepository
	chatgo     ChatGoService
}

func NewAddTask(repository domain.TaskRepository, chatgo ChatGoService) *AddTask {
	return &AddTask{repository: repository, chatgo: chatgo}
}

func (h AddTask) Handle(title string, description string) (uuid.UUID, error) {
	id, err := h.chatgo.AddTask(title, description)
	if err != nil {
		return uuid.Nil, err
	}
	task := domain.NewTask(id, title, description)
	h.repository.Add(task)

	return id, nil
}
