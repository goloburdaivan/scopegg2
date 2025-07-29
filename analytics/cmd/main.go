package main

import (
	"scopegg2-analytics/internal/config"
	"scopegg2-analytics/internal/worker"
)

func main() {
	cfg := config.InitConfig()
	w := worker.NewWorker(cfg)
	w.Run()
}
