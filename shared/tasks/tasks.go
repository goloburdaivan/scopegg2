package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

const (
	DemoUploaded = "demo:uploaded"
)

type DemoUploadedPayload struct {
	Filename   string    `json:"filename"`
	Path       string    `json:"path"`
	UploadedAt time.Time `json:"uploaded_at"`
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
