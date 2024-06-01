package domain

type Workflow struct {
	chiefExecutiveOfficer ChiefExecutiveOfficer
	developer             Developer
	reviewer              Reviewer
	storage               Storage
	steps                 map[State]func(task *Task)
}

type State string

const (
	StateNewRequest         State = "NewRequest"
	StateCodding            State = "Codding"
	StateReview             State = "Review"
	StateReviewModification State = "ReviewModification"
	StateComplete           State = "Complete"
	StateClose              State = "Close"
)

func NewWorkflow(chiefExecutiveOfficer ChiefExecutiveOfficer, developer Developer, reviewer Reviewer, storage Storage) *Workflow {
	steps := map[State]func(project *Task){
		StateNewRequest:         chiefExecutiveOfficer.DescribeTask,
		StateCodding:            developer.Codding,
		StateReview:             reviewer.Review,
		StateReviewModification: developer.ReviewModification,
		StateComplete: func(project *Task) {
			storage.SaveFiles(project.Files)
			project.State = StateClose
		},
	}

	return &Workflow{chiefExecutiveOfficer: chiefExecutiveOfficer, developer: developer, reviewer: reviewer, steps: steps}
}

func (w *Workflow) Do(task *Task) {
	for {
		w.steps[task.State](task)
		if task.State == StateClose {
			break
		}
	}
}
