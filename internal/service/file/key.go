package file

import (
	"fmt"
	"mime/multipart"

	"github.com/fivemanage/lite/internal/crypt"
	"github.com/gabriel-vasile/mimetype"
)

func generateFileKey(organizationID string, file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
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
	key := fmt.Sprintf("%s/%s.%s", organizationID, filename, mime.Extension())

	return key, mime.String(), nil
}
