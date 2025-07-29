package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"scopegg2-analytics/internal/collections"
	"scopegg2-infra/taskqueue/interfaces"
	"scopegg2-shared/tasks"
)

type DemoProcessor interface {
	ProcessDemo(ctx context.Context, demoPath string, steamId uint64) (*collections.Highlights, error)
}

type AnalyticsHandler struct {
	taskQueue     interfaces.TaskQueue
	demoProcessor DemoProcessor
}

func NewAnalyticsHandler(
	taskQueue interfaces.TaskQueue,
	demoProcessor DemoProcessor,
) *AnalyticsHandler {
	return &AnalyticsHandler{
		demoProcessor: demoProcessor,
		taskQueue:     taskQueue,
	}
}

func (handler *AnalyticsHandler) AnalyzeDemo(ctx context.Context, t *asynq.Task) error {
	var event tasks.DemoUploadedPayload
	if err := json.Unmarshal(t.Payload(), &event); err != nil {
		return fmt.Errorf("unmarshal payload: %w", err)
	}

	log.Printf("Processing demo uploaded event...")

	/**
	TODO: FIX FOR Ilya
	SteamID - Хардкод, он должен приходит в событии DemoUploadedPayload, нужен фикс.
	*/
	highlights, err := handler.demoProcessor.ProcessDemo(ctx, event.Path, 76561199184042835)
	if err != nil {
		log.Printf("process demo: %s", err.Error())
		return fmt.Errorf("process demo: %w", err)
	}

	demoAnalyzedTask, err := tasks.NewDemoAnalyzedPayload(event.Path, highlights.GetData())

	err = handler.taskQueue.Enqueue(demoAnalyzedTask)
	if err != nil {
		log.Printf("enqueue task: %s", err.Error())
		return err
	}

	log.Printf("Demo analyzed successfully")

	return nil
}
