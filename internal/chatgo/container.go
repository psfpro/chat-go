package chatgo

import (
	"context"
	"log"

	"chatgo/internal/chatgo/application"
	"chatgo/internal/chatgo/domain"
	"chatgo/internal/chatgo/infrastructure/agent"
	"chatgo/internal/chatgo/infrastructure/agent/openai"
	grpcApi "chatgo/internal/chatgo/infrastructure/api/grpc"
	"chatgo/internal/chatgo/infrastructure/authentication"
	"chatgo/internal/chatgo/infrastructure/persistance/postgres"
	"chatgo/internal/chatgo/infrastructure/storage"
	"chatgo/proto"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
)

type Container struct {
	app *App
}

func (c *Container) App() *App {
	return c.app
}

func NewContainer() *Container {
	ctx := context.Background()
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
	// Repository
	userRepository := postgres.NewUserRepository(db)
	if err = userRepository.CreateTable(ctx); err != nil {
		log.Fatalf("create table error: %v", err)
	}
	// Services
	openaiClient := openai.NewClient(config.openAiApiKey)
	ceo := openai.NewChiefExecutiveOfficer(openaiClient)
	developer := openai.NewDeveloper(openaiClient)
	reviewer := openai.NewReviewer(openaiClient)
	workflow := domain.NewWorkflow(ceo, developer, reviewer, storage.NewFilesystem())
	agentWorker := agent.NewWorker(workflow)
	authenticationService := authentication.NewService(config.jwtPrivateKey, config.jwtPublicKey)
	userLoginHandler := application.NewUserLoginHandler(userRepository, authenticationService)
	userRegistrationHandler := application.NewUserRegistrationHandler(userRepository, authenticationService)
	addTaskHandler := application.NewAddTask(agentWorker)
	srv := grpcApi.NewChatGoServer(authenticationService, userLoginHandler, userRegistrationHandler, addTaskHandler)
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcApi.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcApi.UnaryServerInterceptor()),
	)
	proto.RegisterChatGoServer(grpcServer, srv)

	app := NewApp(agentWorker, grpcServer)

	return &Container{
		app: app,
	}
}
