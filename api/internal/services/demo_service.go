package services

import (
	"context"
	"io"
	"log"
	"scopegg2-infra/taskqueue/interfaces"
	"scopegg2-shared/tasks"
	"scopegg2/internal/dto"
	"scopegg2/internal/handlers"
)

type DemoUploader interface {
	Upload(ctx context.Context, file io.Reader, filename string) (*dto.FileUploadedResult, error)
}

type demoService struct {
	taskQueue interfaces.TaskQueue
	uploader  DemoUploader
}

func NewDemoService(
	taskQueue interfaces.TaskQueue,
	uploader DemoUploader,
) handlers.DemoService {
	return &demoService{
		uploader:  uploader,
		taskQueue: taskQueue,
	}
}

func (s *demoService) UploadAndNotify(ctx context.Context, file io.Reader, filename string) (*dto.FileUploadedResult, error) {
	result, err := s.uploader.Upload(ctx, file, filename)

	if err != nil {
		log.Printf("Failed to upload file: %s\n", err.Error())
		return nil, err
	}

	task, err := tasks.NewDemoUploadedPayload(result.Filename, result.Path, result.UploadedAt)

	if err != nil {
		log.Printf("Failed to create demo uploaded payload: %s\n", err.Error())
		return nil, err
	}

	err = s.taskQueue.Enqueue(task)
	if err != nil {
		log.Printf("Failed to enqueue task: %s\n", err.Error())
		return nil, err
	}

	return result, nil
}
