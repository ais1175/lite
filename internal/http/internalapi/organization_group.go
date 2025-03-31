package internalapi

import (
	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/service/organization"
	"github.com/labstack/echo/v4"
)

func registerOrganizationApi(group *echo.Group, organizationService *organization.Service) {
	group.POST("/organization", func(c echo.Context) error {
		var data api.CreateOrganizationRequest
		if err := c.Bind(&data); err != nil {
			return echo.NewHTTPError(400, err)
		}

		organization, err := organizationService.CreateOrganization(c.Request().Context(), &data)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}

		return c.JSON(200, organization)
	})

	group.GET("/organization", func(c echo.Context) error {
		organizations, err := organizationService.ListOrganizations(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(500, err)
		}

		return c.JSON(200, httputil.Response(organizations))
	})
}
