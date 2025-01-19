package s3

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	client *s3.Client
}

func New() *Storage {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		// TODO: handle error properly
		panic(err)
	}

	// s3 will automatically use the following environment variables:
	// AWS_ACCESS_KEY_ID
	// AWS_SECRET_ACCESS_KEY
	// these are custom:
	// AWS_REGION
	// AWS_ENDPOINT
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = os.Getenv("AWS_REGION")
		o.BaseEndpoint = aws.String(os.Getenv("AWS_ENDPOINT"))
		o.EndpointResolverV2 = s3.NewDefaultEndpointResolverV2()
	})

	return &Storage{
		client: client,
	}
}

// UploadFile will both upload and replace as long as the key is the same
func (s *Storage) UploadFile() error {
	s.client.PutObject(context.TODO(), &s3.PutObjectInput{})
	return nil
}

func (s *Storage) DeleteFile() error {
	s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{})
	return nil
}
