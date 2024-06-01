package openai

import (
	"chatgo/internal/chatgo/domain"
)

type Reviewer struct {
	client         *Client
	role           string
	specialisation string
}

func NewReviewer(client *Client) *Reviewer {
	return &Reviewer{
		client: client,
		role:   `Reviewer`,
		specialisation: `
You can help programmers to assess source codes for software troubleshooting,
fix bugs to increase code quality and robustness, and offer proposals to improve the source codes.
Formulated the following regulations:
1) all referenced classes should be imported;
2) all methods should be implemented;
3) all methods need to have the necessary comments;
4) no potential bugs;
5) The entire project conforms to the tasks proposed by the user;
6) most importantly, do not only check the errors in the code, but also the logic of code. 
`,
	}
}

func (r *Reviewer) Review(project *domain.Task) {
	code := ""
	for _, file := range project.Files {
		code += fileToString(file)
	}
	message := `
Programming language: Go
Codes:
` + code + `
Make sure that user can interact with generated software without losing any feature in the requirement.
Check the above regulations one by one and review the codes in detail, propose one comment with the highest priority about the codes,
and give me instructions on how to fix. 
Comment with the highest priority and corresponding suggestions on revision.
If the codes are perfect and you have no comment on them, return only one line like \"Finished\"."
`
	response := r.client.Chat(r.role, r.specialisation, message)
	project.Comments = response

	if response == "Finished" {
		project.State = domain.StateComplete
	} else {
		project.State = domain.StateReviewModification
	}
}
