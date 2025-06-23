package publicapi

import (
	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/auth"
	"github.com/fivemanage/lite/internal/service/log"
	"github.com/labstack/echo/v4"
)

func registerLogsApi(group *echo.Group, logService *log.Service) {
	group.POST("/batch", func(c echo.Context) error {
		ctx := c.Request().Context()

		orgID, err := auth.CurrentOrgId(c)
		if err != nil {
			return err
		}

		var logs []api.Log
		if err := c.Bind(&logs); err != nil {
			return err
		}

		logService.SubmitLogs(ctx, orgID, "dataset", logs)

		return nil
	})
}
