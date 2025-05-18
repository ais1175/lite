package file

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/internal/database/query/file"
	"github.com/fivemanage/lite/internal/http/httputil"
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
	file multipart.File,
	fileHeader *multipart.FileHeader,
) error {
	var err error
	var key string

	primaryKey, err := crypt.GeneratePrimaryKey()
	if err != nil {
		return UploadStorageError{
			ErrorMsg: err.Error(),
		}
	}

	mimeType, ext, fileType, err := httputil.GetMimeDetails(fileHeader, file)
	if err != nil {
		return UploadStorageError{
			ErrorMsg: errors.New("failed to get mime type").Error(),
		}
	}

	key, err = generateFileKey(organizationID, ext)
	if err != nil {
		return UploadStorageError{
			ErrorMsg: errors.New("failed to generate file key").Error(),
		}
	}

	// add organizationID to the asset
	asset := &database.Asset{
		ID:   primaryKey,
		Type: fileType,
		Size: fileHeader.Size,
		Key:  key,
	}

	tx, err := filequery.Create(ctx, s.db, asset)
	if err != nil {
		return err
	}

	err = s.storage.UploadFile(ctx, file, key, mimeType)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	// this is a bit tricky, but if this fails....then...oh well
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateStorageFile(
	ctx context.Context,
	organizationID string,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) error {
	var err error

	primaryKey, err := crypt.GeneratePrimaryKey()
	if err != nil {
		return UploadStorageError{
			ErrorMsg: err.Error(),
		}
	}

	mimeType, _, fileType, err := httputil.GetMimeDetails(fileHeader, file)
	if err != nil {
		return UploadStorageError{
			ErrorMsg: errors.New("failed to get mime type").Error(),
		}
	}

	// fileHeader.Filename has the extension most of the time
	// should it become an issue, we can look for it and check if its empty;
	key := generateWebKey(organizationID, fileHeader.Filename)

	// add organizationID to the asset
	asset := &database.Asset{
		ID:             primaryKey,
		OrganizationID: organizationID,
		Type:           fileType,
		Size:           fileHeader.Size,
		Key:            key,
	}

	tx, err := filequery.Create(ctx, s.db, asset)
	if err != nil {
		return err
	}

	err = s.storage.UploadFile(ctx, file, key, mimeType)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ListStorageFiles(
	ctx context.Context,
	organizationID string,
	search string,
) (*api.AssetResponse, error) {
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
			ID:        file.ID,
			Type:      file.Type,
			Key:       file.Key,
			Size:      file.Size,
			CreatedAt: file.CreatedAt,
		}
	}

	totalCount, err := filequery.FindTotalStorageCount(ctx, s.db, organizationID)
	if err != nil {
		storageError := &ListStorageError{
			ErrorMsg: err.Error(),
		}

		logrus.WithError(storageError).
			WithField("organization_id", organizationID).Error("FileService.ListStorageFiles")

		return nil, storageError
	}

	response := &api.AssetResponse{
		StorageFiles: assets,
		TotalCount:   totalCount,
	}

	return response, nil
}
