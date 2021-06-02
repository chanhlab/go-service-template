package workers

import (
	"fmt"

	JobWorkers "github.com/digitalocean/go-workers2"
)

// ExampleWorker ...
type ExampleWorker struct {
	Manager     *JobWorkers.Manager
	queue       string
	concurrency int
}

// NewExampleWorker create instance
func NewExampleWorker(manager *JobWorkers.Manager, queue string, concurrency int) *ExampleWorker {
	return &ExampleWorker{
		Manager:     manager,
		queue:       queue,
		concurrency: concurrency,
	}
}

// Register is a registration worker
func (worker *ExampleWorker) Register() {
	workerMiddleware := NewWorkerMiddleware()
	worker.Manager.AddWorker(worker.queue, worker.concurrency, worker.Execute, workerMiddleware...)
}

// Execute is a execution job of worker
func (worker *ExampleWorker) Execute(message *JobWorkers.Msg) error {
	fmt.Printf("\nJob ID: %s", message.Jid())
	fmt.Printf("\nID: %v\n", message.Args().Interface())

	return nil
}
