package httputil

import (
	"mime/multipart"
	"net/http"
)

func File(r *http.Request, key string) (file multipart.File, fileHeader *multipart.FileHeader, err error) {
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, nil, err
	}

	// check if "file" exists in the form first
	if r.MultipartForm.File["file"] != nil {
		file, fileHeader, err = r.FormFile("file")
		return file, fileHeader, err
	}

	file, fileHeader, err = r.FormFile(key)
	return file, fileHeader, err
}

func Metadata(r *http.Request) string {
	return r.FormValue("metadata")
}

func Resource(r *http.Request) string {
	return r.FormValue("resource")
}

func Filename(r *http.Request) string {
	return r.FormValue("filename")
}
