package internalapi

import (
	"errors"
	"net/http"

	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func registerStorageApi(group *echo.Group, fileService *file.Service) {
	group.GET("/storage/:organizationId", func(c echo.Context) error {
		ctx := c.Request().Context()

		organizationID := c.Param("organizationId")
		search := c.QueryParam("search")
		fileType := c.QueryParam("type")

		page, _ := httputil.QueryInt(c, "page", 0)
		pageSize, _ := httputil.QueryInt(c, "pageSize", 20)

		assetData, err := fileService.ListStorageFiles(ctx, organizationID, search, fileType, page, pageSize)
		if err != nil {
			if errors.Is(err, file.ListStorageError{}) {
				return echo.NewHTTPError(http.StatusInternalServerError,
					httputil.ErrorResponse("Failed to list storage files"),
				)
			}

			return echo.NewHTTPError(http.StatusInternalServerError, httputil.ErrorResponse(err.Error()))
		}

		return c.JSON(200, httputil.Response(assetData))
	})

	group.GET("/storage/:organizationId/file/:fileId", func(c echo.Context) error {
		ctx := c.Request().Context()

		organizationID := c.Param("organizationId")
		fileID := c.Param("fileId")

		assetData, err := fileService.GetStorageFile(ctx, organizationID, fileID)
		if err != nil {
			if errors.Is(err, file.GetFileError{}) {
				return echo.NewHTTPError(http.StatusInternalServerError,
					httputil.ErrorResponse("Failed to get storage file"),
				)
			}

			return echo.NewHTTPError(http.StatusInternalServerError, httputil.ErrorResponse(err.Error()))
		}

		return c.JSON(200, httputil.Response(assetData))
	})

	group.POST("/storage/:organizationId/upload", func(c echo.Context) error {
		ctx := c.Request().Context()

		organizationID := c.Param("organizationId")

		formFile, header, err := httputil.File(c.Request(), "file")
		if err != nil {
			logrus.WithField("organizationId", organizationID).Error("failed to get file from request")

			return echo.NewHTTPError(http.StatusInternalServerError,
				httputil.ErrorResponse("Failed to get file from request"),
			)
		}

		err = fileService.CreateStorageFile(
			ctx,
			organizationID,
			formFile,
			header,
		)
		if err != nil {
			logrus.WithError(err).WithField("organizationId", organizationID).Error("failed to create file")

			if errors.Is(err, file.UploadStorageError{}) {
				return echo.NewHTTPError(http.StatusInternalServerError,
					httputil.ErrorResponse("Failed to upload storage file"),
				)
			}

			return echo.NewHTTPError(http.StatusInternalServerError, httputil.ErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, httputil.Response("File uploaded successfully"))
	})
}
