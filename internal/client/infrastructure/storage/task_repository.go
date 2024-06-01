package storage

import "chatgo/internal/client/domain"

type TaskRepository struct {
	data map[string]*domain.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{data: make(map[string]*domain.Task)}
}

func (r *TaskRepository) GetAll() map[string]*domain.Task {
	return r.data
}

func (r *TaskRepository) Add(task *domain.Task) {
	r.data[task.ID.String()] = task
}
