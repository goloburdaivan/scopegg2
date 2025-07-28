package services

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"scopegg2/internal/config"
	"scopegg2/internal/dto"
	"time"
)

type s3DemoUploader struct {
	client *s3.Client
	cfg    *config.Config
}

func NewS3DemoUploader(awsConfig aws.Config, cfg *config.Config) DemoUploader {
	client := s3.NewFromConfig(awsConfig)

	return &s3DemoUploader{client: client, cfg: cfg}
}

func (s *s3DemoUploader) Upload(ctx context.Context, file io.Reader, filename string) (*dto.FileUploadedResult, error) {
	objectKey := fmt.Sprintf("demos/%s", filename)

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.cfg.BucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})

	if err != nil {
		log.Printf("Failed to upload file to S3: %s\n", err.Error())
		return &dto.FileUploadedResult{}, err
	}

	return &dto.FileUploadedResult{
		Filename:   filename,
		Path:       objectKey,
		UploadedAt: time.Now(),
	}, nil
}
