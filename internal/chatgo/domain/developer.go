package domain

type Developer interface {
	Codding(task *Task)
	ReviewModification(task *Task)
}
