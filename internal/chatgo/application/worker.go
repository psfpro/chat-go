package application

import "chatgo/internal/chatgo/domain"

type Worker interface {
	Do(task *domain.Task)
}
