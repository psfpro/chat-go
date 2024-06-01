package openai

import (
	"chatgo/internal/chatgo/domain"
	"regexp"
	"strings"
)

type Developer struct {
	client         *Client
	role           string
	specialisation string
}

func NewDeveloper(client *Client) *Developer {
	return &Developer{
		client: client,
		role:   `Developer`,
		specialisation: `
You are Go developer. Use all features from go version 1.22.
Use next response format example:
main.go
` + "```" + `go
		package main

		import "fmt"

		func main() {
		fmt.Println("Hello, world!")
	}
` + "```" + `
		`,
	}
}

func (d *Developer) Codding(task *domain.Task) {
	message := `
You will first lay out the names of the core classes, functions, methods that will be necessary, as well as a quick comment on their purpose.
Then you will output the content of each file including complete code. Each file must strictly follow a markdown code block format,
where the following tokens must be replaced such that \"FILENAME\" is the lowercase file name including the file extension,
"PATH" is the path to the file relative to the root of the task, \"LANGUAGE\" in the programming language,
\"DOCSTRING\" is a string literal specified in source code that is used to document a specific segment of code, and \"CODE\" is the original code:
PATH/FILENAME
` + "```LANGUAGE" + `
CODE
` + "```" + `
You will start with the \"main\" file, then go to the ones that are imported by that file, and so on.
Please note that the code should be fully functional. Ensure to implement all functions.
Don't forget to add the go.mod file
`
	response := d.client.Chat(d.role, d.specialisation, message+task.TechnicalSolution)

	files := newFiles(response)
	for path, file := range files {
		task.Files[path] = file
	}
	task.State = domain.StateReview
}

func (d *Developer) ReviewModification(task *domain.Task) {
	code := ""
	for _, file := range task.Files {
		code += fileToString(file)
	}
	message := `
You will output the content of each file including complete code. Each file must strictly follow a markdown code block format,
where the following tokens must be replaced such that \"FILENAME\" is the lowercase file name including the file extension,
"PATH" is the path to the file relative to the root of the task, \"LANGUAGE\" in the programming language,
\"DOCSTRING\" is a string literal specified in source code that is used to document a specific segment of code, and \"CODE\" is the original code:
PATH/FILENAME
` + "```LANGUAGE" + `
CODE
` + "```" + `

Codes:` + code + `
Code comments:` + task.Comments + `
You should modify corresponding codes according to the comments. Then, output the full and complete codes with all bugs fixed based on the comments. Return all codes strictly following the required format.
`
	response := d.client.Chat(d.role, d.specialisation, message)

	files := newFiles(response)
	for path, file := range files {
		task.Files[path] = file
	}
	if task.Iteration > 1 {
		task.State = domain.StateComplete
	} else {
		task.Iteration++
		task.State = domain.StateReview
	}
}

func newFiles(text string) map[string]*domain.File {
	result := make(map[string]*domain.File)
	re := regexp.MustCompile("(.+?)\\n```(.+?)\\n([\\s\\S]*?)```")
	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		path := match[1]
		path = strings.Trim(path, "#*` ")
		result[path] = &domain.File{
			Path:     path,
			Language: match[2],
			Content:  match[3],
		}
	}
	return result
}

func fileToString(f *domain.File) string {
	return f.Path + "\n```" + f.Language + "\n" + f.Content + "```\n"
}
