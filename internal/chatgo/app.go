package chatgo

import (
	"chatgo/internal/chatgo/infrastructure/agent"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	agentWorker *agent.Worker
}

func NewApp(agentWorker *agent.Worker) *App {
	return &App{agentWorker: agentWorker}
}

func (a *App) Run() {
	a.agentWorker.Start()
	a.waitStopSignal()
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
}
