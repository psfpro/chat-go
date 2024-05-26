package client

import "log"

type Container struct {
	app *App
}

func (c *Container) App() *App {
	return c.app
}

func NewContainer() *Container {
	config := NewConfig()
	log.Printf("server address %v", config.serverAddress)

	app := NewApp()

	return &Container{
		app: app,
	}
}
