package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"scopegg2-shared/dto"
	"time"
)

const (
	DemoUploaded = "demo:uploaded"
	DemoAnalyzed = "demo:analyzed"
)

type DemoUploadedPayload struct {
	Filename   string    `json:"filename"`
	Path       string    `json:"path"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type HighlightsPayload struct {
	DemoPath   string             `json:"demo_path"`
	Highlights map[int][]dto.Kill `json:"highlights"`
}

func NewDemoUploadedPayload(filename string, path string, uploadedAt time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(DemoUploadedPayload{
		Filename:   filename,
		Path:       path,
		UploadedAt: uploadedAt,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(DemoUploaded, payload), nil
}

func NewDemoAnalyzedPayload(demoPath string, highlights map[int][]dto.Kill) (*asynq.Task, error) {
	payload, err := json.Marshal(HighlightsPayload{
		DemoPath:   demoPath,
		Highlights: highlights,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(DemoAnalyzed, payload), nil
}
