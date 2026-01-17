package internalapi

import (
	"net/http"

	"github.com/fivemanage/lite/internal/service/system"
	"github.com/labstack/echo/v4"
)

func registerSystemApi(group *echo.Group, systemService *system.Service) {
	group.GET("/system/version", func(c echo.Context) error {
		status, err := systemService.GetVersionStatus()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, status)
	})
}
