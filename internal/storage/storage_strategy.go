package storage

import (
	"context"
	"io"

	"github.com/fivemanage/lite/internal/storage/s3"
)

type StorageLayer interface {
	UploadFile(context.Context, io.Reader, string, string) error
	DeleteFile() error
}

func New(provider string) StorageLayer {
	switch provider {
	case "s3":
		return s3.New()
	}

	return nil
}
