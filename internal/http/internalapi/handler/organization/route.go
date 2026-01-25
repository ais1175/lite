package organization

import (
	"github.com/fivemanage/lite/internal/service/organization"
	"github.com/labstack/echo/v4"
)

type handler struct {
	organizationService *organization.Service
}

func RegisterRoutes(group *echo.Group, organizationService *organization.Service) {
	h := handler{
		organizationService: organizationService,
	}

	group.POST("/organization", h.createOrganizationHandler)
	group.GET("/organization", h.listOrganizationsHandler)
	group.GET("/organization/:id", h.getOrganizationHandler)
	group.GET("/organization/:id/stats", h.getOrganizationStatsHandler)
}
