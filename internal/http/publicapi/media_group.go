package publicapi

import (
	"net/http"

	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/labstack/echo/v4"
)

func registerMediaApi(group *echo.Group, fileService *file.Service) {
	group.POST("/image", func(c echo.Context) error {
		var err error
		ctx := c.Request().Context()

		file, header, err := httputil.File(c.Request(), "image")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		err = fileService.CreateFile(ctx, "image", file, header)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.JSON(200, nil)
	})
}
