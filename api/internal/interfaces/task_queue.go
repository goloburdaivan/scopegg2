package interfaces

import "github.com/hibiken/asynq"

type TaskQueue interface {
	Enqueue(task *asynq.Task) error
}
