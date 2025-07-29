package redis

import (
	"github.com/hibiken/asynq"
	"scopegg2-shared/interfaces"
)

type redisTaskQueue struct {
	client *asynq.Client
}

func NewRedisTaskQueue(client *asynq.Client) interfaces.TaskQueue {
	return &redisTaskQueue{client: client}
}

func (r redisTaskQueue) Enqueue(task *asynq.Task) error {
	_, err := r.client.Enqueue(task)

	return err
}
