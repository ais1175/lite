package file

import (
	"context"
	"mime/multipart"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/internal/database/query/file"
	"github.com/fivemanage/lite/internal/storage"
	"github.com/sirupsen/logrus"
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
	organizationID string,
	fileType string,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) error {
	var err error
	key, contentType, err := generateFileKey(organizationID, file, fileHeader)
	if err != nil {
		return err
	}

	asset := &database.Asset{
		Type: fileType,
		Size: fileHeader.Size,
		Key:  key,
	}

	tx, err := filequery.Create(ctx, s.db, asset)
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

func (s *Service) ListStorageFiles(
	ctx context.Context,
	organizationID string,
	search string,
) ([]*api.Asset, error) {
	var err error

	files, err := filequery.FindStorageFiles(ctx, s.db, organizationID, search)
	if err != nil {
		storageError := &ListStorageError{
			ErrorMsg: err.Error(),
		}
		logrus.WithError(storageError).
			WithField("organization_id", organizationID).
			Error("FileService.ListStorageFiles")
		return nil, storageError
	}

	assets := make([]*api.Asset, len(files))
	for i, file := range files {
		assets[i] = &api.Asset{
			ID:   file.ID,
			Type: file.Type,
			Key:  file.Key,
			Size: file.Size,
		}
	}

