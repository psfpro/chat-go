package openai

import (
	"chatgo/internal/chatgo/domain"
)

type ChiefExecutiveOfficer struct {
	client         *Client
	role           string
	specialisation string
}

func NewChiefExecutiveOfficer(client *Client) *ChiefExecutiveOfficer {
	return &ChiefExecutiveOfficer{
		client: client,
		role:   "Chief Executive Officer",
		specialisation: `
Your main responsibilities include being an active decision-maker on users' demands and other key policy issues, leader, manager, and executor.
Your decision-making role involves high-level decisions about policy and strategy; 
and your communicator role can involve speaking to the organization's management and employees.`,
	}
}

func (o *ChiefExecutiveOfficer) DescribeTask(task *domain.Task) {
	message := `We have short description of a software design requirement, 
please rewrite it into a detailed prompt that can make large language model know how to make this software better based this prompt,
the prompt should ensure LLMs build a software that can be run correctly, which is the most import part you need to consider.
remember that the revised prompt should not contain more than 200 words, 
here is the short description:"` + task.Title + `".`

	response := o.client.Chat(o.role, o.specialisation, message)
	task.TechnicalSolution = response
	task.State = domain.StateCodding
}
