//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/google/wire"
	"github.com/hibiken/asynq"
	config2 "scopegg2/internal/config"
	"scopegg2/internal/handlers"
	"scopegg2/internal/queue/redis"
	server2 "scopegg2/internal/server"
	services2 "scopegg2/internal/services"
)

func InitializeApp(cfg *config2.Config) (*server2.App, error) {
	wire.Build(
		newRedisClient,
		redis.NewRedisTaskQueue,
		newAWSConfig,
		services2.NewS3DemoUploader,
		services2.NewDemoService,
		handlers.NewUploadDemoHandler,
		server2.NewRouter,
		server2.NewApp,
	)
	return &server2.App{}, nil
}

func newRedisClient(cfg *config2.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.RedisUrl,
	})
}

func newAWSConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO())
}
