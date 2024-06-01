package tui

import (
	"chatgo/internal/client/application"
	tea "github.com/charmbracelet/bubbletea"
)

type Program struct {
	teaProgram *tea.Program
}

func NewProgram(addTaskHandler *application.AddTask) *Program {
	return &Program{
		teaProgram: tea.NewProgram(initialModel(addTaskHandler), tea.WithoutSignalHandler()),
	}
}

func (p Program) Run() error {
	_, err := p.teaProgram.Run()
	if err != nil {
		return err
	}
	return nil
}
