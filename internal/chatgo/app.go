package chatgo

import (
	"chatgo/internal/chatgo/infrastructure/agent"
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	agentWorker *agent.Worker
	grpcServer  *grpc.Server
}

func NewApp(agentWorker *agent.Worker, server *grpc.Server) *App {
	return &App{agentWorker: agentWorker, grpcServer: server}
}

func (a *App) Run() {
	a.agentWorker.Start()
	a.runGrpcServer()
	a.waitStopSignal()
}

func (a *App) runGrpcServer() {
	go func() {
		log.Println("Starting gRPC server")
		listen, err := net.Listen("tcp", ":3200")
		if err != nil {
			log.Fatal(err)
		}
		if err := a.grpcServer.Serve(listen); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()
}

func (a *App) waitStopSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-signalChan
	signal.Stop(signalChan)
	log.Printf("received signal %s, shutting down", sig.String())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.shutdown(ctx)
}

func (a *App) shutdown(_ context.Context) {
	a.agentWorker.Stop()
	a.grpcServer.GracefulStop()
}
