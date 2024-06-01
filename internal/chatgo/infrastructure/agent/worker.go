package agent

import (
	"chatgo/internal/chatgo/domain"
	"context"
	"log"
	"sync"
)

type Worker struct {
	tasks    chan *domain.Task
	wg       sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
	workflow *domain.Workflow
}

func NewWorker(workflow *domain.Workflow) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		tasks:    make(chan *domain.Task, 1),
		ctx:      ctx,
		cancel:   cancel,
		workflow: workflow,
	}
}

func (w *Worker) Do(task *domain.Task) {
	w.tasks <- task
}

func (w *Worker) Start() {
	log.Println("Starting worker")
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			select {
			case task := <-w.tasks:
				log.Println("Starting task")
				w.workflow.Do(task)
				log.Println("Finished task")
			case <-w.ctx.Done():
				log.Println("Worker received shutdown signal")
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.cancel()
	w.wg.Wait()
	log.Println("Worker stopped")
}
