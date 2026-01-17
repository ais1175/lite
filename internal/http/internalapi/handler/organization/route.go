package organization

import (
	"github.com/fivemanage/lite/internal/service/organization"
	"github.com/labstack/echo/v4"
)

type handler struct {
	organizationService *organization.Service
}

func RegisterRoutes(group *echo.Group, organizationService *organization.Service) {
	handler := handler{
		organizationService: organizationService,
	}

	group.POST("/organization", handler.createOrganizationHandler)
	group.GET("/organization", handler.listOrganizationsHandler)
	group.GET("/organization/:id", handler.getOrganizationHandler)
}
