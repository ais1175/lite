package publicapi

import (
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/labstack/echo/v4"
)

func Add(group *echo.Group, fileService *file.Service) {
	registerMediaApi(group, fileService)
}
