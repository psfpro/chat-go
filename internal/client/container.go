package client

import (
	"chatgo/internal/client/application"
	"chatgo/internal/client/infrastructure/chatgo/dumb"
	"chatgo/internal/client/infrastructure/storage"
	"chatgo/internal/client/infrastructure/tui"
	"log"
)

type Container struct {
	app *App
}

func (c *Container) App() *App {
	return c.app
}

func NewContainer() *Container {
	config := NewConfig()
	log.Printf("server address %v", config.serverAddress)
	// Repositories
	taskRepository := storage.NewTaskRepository()
	// Services
	chatgo := dumb.ChatGoService{}
	addTaskHandler := application.NewAddTask(taskRepository, chatgo)
	prog := tui.NewProgram(addTaskHandler)

	app := NewApp(prog)

	return &Container{
		app: app,
	}
}
