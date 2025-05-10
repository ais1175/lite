package api

import "mime/multipart"

type UploadFile struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
}

type Asset struct {
	ID   string
	Key  string
	Size int64
	Type string
}

type AssetResponse struct {
	StorageFiles []Asset `json:"storageFiles"`
	Total        int64   `json:"total"`
}
