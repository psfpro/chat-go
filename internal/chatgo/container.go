package chatgo

import (
	"chatgo/internal/chatgo/application"
	"chatgo/internal/chatgo/domain"
	"chatgo/internal/chatgo/infrastructure/agent"
	"chatgo/internal/chatgo/infrastructure/agent/openai"
	grpcApi "chatgo/internal/chatgo/infrastructure/api/grpc"
	"chatgo/internal/chatgo/infrastructure/storage"
	"chatgo/proto"
	"database/sql"
	"google.golang.org/grpc"
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
	// Services
	openaiClient := openai.NewClient(config.openAiApiKey)
	ceo := openai.NewChiefExecutiveOfficer(openaiClient)
	developer := openai.NewDeveloper(openaiClient)
	reviewer := openai.NewReviewer(openaiClient)
	workflow := domain.NewWorkflow(ceo, developer, reviewer, storage.NewFilesystem())
	agentWorker := agent.NewWorker(workflow)
	addTaskHandler := application.NewAddTask(agentWorker)
	srv := grpcApi.NewChatGoServer(addTaskHandler)
	grpcServer := grpc.NewServer()
	proto.RegisterChatGoServer(grpcServer, srv)

	app := NewApp(agentWorker, grpcServer)

	return &Container{
		app: app,
	}
}
