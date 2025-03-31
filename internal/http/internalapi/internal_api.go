package internalapi

import (
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/fivemanage/lite/internal/service/organization"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/labstack/echo/v4"
)

func Add(group *echo.Group, authService *auth.Auth, tokenService *token.Service, organizationService *organization.Service) {
	registerAuthApi(group, authService)
	registerTokensApi(group, tokenService)
	registerOrganizationApi(group, organizationService)
}
