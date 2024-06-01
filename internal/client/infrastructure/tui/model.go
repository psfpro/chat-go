package tui

import (
	"chatgo/internal/client/application"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gofrs/uuid"
	"log"
)

type task struct {
	id          uuid.UUID
	title       string
	description string
}

func (t task) Title() string       { return t.title }
func (t task) Description() string { return t.description }
func (t task) FilterValue() string { return t.title }

type model struct {
	list           list.Model
	titleInput     textinput.Model
	descInput      textinput.Model
	adding         bool
	addTaskHandler *application.AddTask
}

func initialModel(addTaskHandler *application.AddTask) model {
	m := model{
		list:           list.New([]list.Item{}, list.NewDefaultDelegate(), 50, 20),
		titleInput:     textinput.New(),
		descInput:      textinput.New(),
		adding:         false,
		addTaskHandler: addTaskHandler,
	}
	m.titleInput.Placeholder = "New task title"
	m.descInput.Placeholder = "New task description"
	m.list.Title = "ChatGo"
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.adding {
			switch msg.String() {
			case "ctrl+c":
				m.titleInput.Reset()
				m.descInput.Reset()
				m.adding = false
				return m, nil
			case "enter":
				if m.adding {
					if m.titleInput.Focused() {
						m.titleInput.Blur()
						m.descInput.Focus()
						return m, nil
					}
					newTask := task{title: m.titleInput.Value(), description: m.descInput.Value()}
					id, err := m.addTaskHandler.Handle(newTask.title, newTask.description)
					if err != nil {
						log.Println("add task error:", err)
					}
					newTask.id = id
					m.list.InsertItem(len(m.list.Items())-1, newTask)
					m.titleInput.Reset()
					m.descInput.Reset()
					m.adding = false
				}
				return m, nil
			}
		} else {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "a":
				m.adding = true
				m.titleInput.Focus()
				return m, nil
			}
		}
	}

	if m.adding {
		if m.titleInput.Focused() {
			m.titleInput, cmd = m.titleInput.Update(msg)
		} else {
			m.descInput, cmd = m.descInput.Update(msg)
		}
		return m, cmd
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.adding {
		return fmt.Sprintf(
			"Add a new task:\n\n%s\n\n%s\n\n%s",
			m.titleInput.View(),
			m.descInput.View(),
			"(enter to save, ctrl+c to quit)",
		)
	}
	return fmt.Sprintf(
		"%s\n\n%s",
		m.list.View(),
		"(a to add a task, q to quit)",
	)
}
