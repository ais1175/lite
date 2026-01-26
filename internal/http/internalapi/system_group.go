package internalapi

import (
	"net/http"

	"github.com/fivemanage/lite/internal/service/system"
	"github.com/labstack/echo/v4"
)

func registerSystemApi(group *echo.Group, systemService *system.Service) {
	h := &systemHandler{systemService: systemService}
	group.GET("/system/version", h.getVersion)
	group.GET("/system/config", h.getConfig)
}

type systemHandler struct {
	systemService *system.Service
}

// getVersion godoc
// @Summary      Get system version
// @Description  Get the current system version and update status
// @Tags         system
// @Produce      json
// @Success      200  {object}  system.VersionStatus
// @Failure      500  {object}  echo.Map
// @Router       /dash/system/version [get]
func (h *systemHandler) getVersion(c echo.Context) error {
	status, err := h.systemService.GetVersionStatus()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, status)
}

// getConfig godoc
// @Summary      Get system config
// @Description  Get the current system configuration
// @Tags         system
// @Produce      json
// @Success      200  {object}  system.SystemConfig
// @Router       /dash/system/config [get]
func (h *systemHandler) getConfig(c echo.Context) error {
	return c.JSON(http.StatusOK, h.systemService.GetConfig())
}
