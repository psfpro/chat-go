package chatgo

import (
	"chatgo/internal/chatgo/domain"
	"chatgo/internal/chatgo/infrastructure/agent"
	"chatgo/internal/chatgo/infrastructure/agent/openai"
	"chatgo/internal/chatgo/infrastructure/storage"
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

	openaiClient := openai.NewClient(config.openAiApiKey)
	ceo := openai.NewChiefExecutiveOfficer(openaiClient)
	developer := openai.NewDeveloper(openaiClient)
	reviewer := openai.NewReviewer(openaiClient)
	workflow := domain.NewWorkflow(ceo, developer, reviewer, storage.NewFilesystem())
	agentWorker := agent.NewWorker(workflow)

	//agentWorker.Do(domain.NewTask("Design a simple calculator console application"))

	app := NewApp(agentWorker)

	return &Container{
		app: app,
	}
}
