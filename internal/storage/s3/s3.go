package s3

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	client *s3.Client
	bucket string
}

func New() *Storage {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
	)
	if err != nil {
		// TODO: handle error properly
		panic(err)
	}

	// s3 will automatically use the following environment variables:
	// AWS_ACCESS_KEY_ID
	// AWS_SECRET_ACCESS_KEY
	// these are custom:
	// AWS_ENDPOINT
	// AWS_BUCKET
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(os.Getenv("AWS_ENDPOINT"))
	})

	bucket := os.Getenv("AWS_BUCKET")
	if bucket == "" {
		logrus.Fatalln("environment variable: AWS_BUCKET is not set")
	}

	return &Storage{
		client: client,
		bucket: bucket,
	}
}

// UploadFile will both upload and replace as long as the key is the same
func (s *Storage) UploadFile(ctx context.Context, file io.Reader, key, contenType string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contenType),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteFile() error {
	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{})
	if err != nil {
		return err
	}

	return nil
}
