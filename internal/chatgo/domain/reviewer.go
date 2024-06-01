package domain

type Reviewer interface {
	Review(project *Task)
}
