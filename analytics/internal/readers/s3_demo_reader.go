package readers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"scopegg2-analytics/internal/config"
	"scopegg2-analytics/internal/services"
)

type s3DemoReader struct {
	s3Client *s3.Client
	cfg      *config.Config
}

func NewS3DemoReader(s3Client *s3.Client, cfg *config.Config) services.DemoReader {
	return &s3DemoReader{
		s3Client: s3Client,
		cfg:      cfg,
	}
}

func (s s3DemoReader) ReadDemo(ctx context.Context, demoPath string) (io.ReadCloser, error) {
	out, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.cfg.BucketName),
		Key:    aws.String(demoPath),
	})

	if err != nil {
		return nil, fmt.Errorf("s3 GetObject: %w", err)
	}

	return out.Body, nil
}
