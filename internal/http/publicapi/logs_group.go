package publicapi

import (
	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/service/log"
	"github.com/labstack/echo/v4"
)

func registerLogsApi(group *echo.Group, logService *log.Service) {
	group.POST("/batch", func(c echo.Context) error {
		var logs []api.Log
		if err := c.Bind(&logs); err != nil {
			return err
		}

		logService.SubmitLogs(logs)

		return nil
	})
}
