package chatgo

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Container struct {
	app *App
}

func (c *Container) App() *App {
	return c.app
}

func NewContainer() *Container {
	config := NewConfig()
	// DB connection
	db, err := sql.Open("pgx", config.dsn)
	if err != nil {
		log.Fatalf("db open error: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("db connection error: %v", err)
	} else {
		log.Printf("db connection ok")
	}
	if config.openAiApiKey == "" {
		log.Fatal("empty OpenAI API key")
	}

	app := NewApp()

	return &Container{
		app: app,
	}
}
