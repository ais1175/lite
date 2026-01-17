package internalapi

import (
	tokenhandler "github.com/fivemanage/lite/internal/http/internalapi/handler/token"
	"github.com/fivemanage/lite/internal/http/middleware"
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/fivemanage/lite/internal/service/dataset"
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/fivemanage/lite/internal/service/organization"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/labstack/echo/v4"
)

func Add(
	group *echo.Group,
	authService *auth.Auth,
	tokenService *token.Service,
	organizationService *organization.Service,
	fileService *file.Service,
	datasetService *dataset.Service,
) {
	group.Use(middleware.Session(authService))

	registerAuthApi(group, authService)
	tokenhandler.RegisterRoutes(group, tokenService)
	registerOrganizationApi(group, organizationService)
	registerStorageApi(group, fileService)
	registerDatasetApi(group, datasetService)
}
