package worker

import (
	"github.com/hibiken/asynq"
	"log"
	"scopegg2-analytics/di"
	"scopegg2-analytics/internal/config"
	"scopegg2-shared/tasks"
)

type Worker struct {
	server *asynq.Server
	mux    *asynq.ServeMux
}

func NewWorker(cfg *config.Config) *Worker {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.RedisUrl},
		asynq.Config{
			Concurrency: 10,
		},
	)

	mux := asynq.NewServeMux()

	return &Worker{server: server, mux: mux}
}

func (worker *Worker) registerHandlers() {
	demoUploadedHandler, err := di.InitializeAnalyticsHandler()
	if err != nil {
		log.Fatal(err)
	}
	worker.mux.HandleFunc(tasks.DemoUploaded, demoUploadedHandler.AnalyzeDemo)
}

func (worker *Worker) Run() {
	worker.registerHandlers()
	if err := worker.server.Run(worker.mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
