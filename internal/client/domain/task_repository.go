package domain

type TaskRepository interface {
	GetAll() map[string]*Task
	Add(task *Task)
}
