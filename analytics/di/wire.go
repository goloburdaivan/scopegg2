//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/wire"
	"github.com/hibiken/asynq"
	"scopegg2-infra/taskqueue/queue/redis"

	config2 "scopegg2-analytics/internal/config"
	"scopegg2-analytics/internal/handlers"
	"scopegg2-analytics/internal/readers"
	"scopegg2-analytics/internal/services"
)

func InitializeAnalyticsHandler() (*handlers.AnalyticsHandler, error) {
	wire.Build(
		config2.InitConfig,
		newAWSConfig,
		newS3Client,
		readers.NewS3DemoReader,
		services.NewDemoInfoCsDemoProcessor,
		newRedisClient,
		redis.NewRedisTaskQueue,
		handlers.NewAnalyticsHandler,
	)
	return &handlers.AnalyticsHandler{}, nil
}

func newRedisClient(cfg *config2.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.RedisUrl,
	})
}

func newAWSConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO())
}

func newS3Client(cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg)
}
