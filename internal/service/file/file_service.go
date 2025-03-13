package file

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
	filequery "github.com/fivemanage/lite/internal/database/query/file"
	"github.com/fivemanage/lite/internal/storage"
	"github.com/gabriel-vasile/mimetype"
	"github.com/uptrace/bun"
)

type Service struct {
	db      *bun.DB
	storage storage.StorageLayer
}

func NewService(db *bun.DB, storageLayer storage.StorageLayer) *Service {
	return &Service{
		db:      db,
		storage: storageLayer,
	}
}

func (s *Service) CreateFile(
	ctx context.Context,
	fileType string,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) error {
	var err error
	key, contentType, err := generateFileKey(fileType, file, fileHeader)
	if err != nil {
		return err
	}

	dbFile := &database.File{
		Type: fileType,
		Size: fileHeader.Size,
		Key:  key,
	}

	tx, err := filequery.Create(ctx, s.db, dbFile)
	if err != nil {
		return err
	}

	err = s.storage.UploadFile(ctx, file, key, contentType)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	return nil
}

func generateFileKey(fileType string, file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
	filename, err := crypt.GenerateFilename()
	if err != nil {
		return "", "", err
	}

	buf := make([]byte, fileHeader.Size)
	_, err = file.Read(buf)
	if err != nil {
		return "", "", err
	}

	mime := mimetype.Detect(buf)
	key := fmt.Sprintf("%s/%s.%s", fileType, filename, mime.Extension())

	return key, mime.String(), nil
}
