package internalapi

import (
	"errors"
	"net/http"

	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/labstack/echo/v4"
)

func registerStorageApi(group *echo.Group, fileService *file.Service) {
	group.GET("/storage/:organizationId", func(c echo.Context) error {
		ctx := c.Request().Context()

		organizationID := c.Param("organizationId")
		search := c.QueryParam("search")

		files, err := fileService.ListStorageFiles(ctx, organizationID, search)
		if err != nil {
			if errors.Is(err, file.ListStorageError{}) {
				return echo.NewHTTPError(http.StatusInternalServerError,
					httputil.ErrorResponse("Failed to list storage files"),
				)
			}

			return echo.NewHTTPError(http.StatusInternalServerError, httputil.ErrorResponse(err.Error()))
		}

		return c.JSON(200, httputil.Response(files))
	})
}
