package token

import (
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/labstack/echo/v4"
)

type handler struct {
	tokenService *token.Service
}

func RegisterRoutes(group *echo.Group, tokenService *token.Service) {
	handler := handler{
		tokenService: tokenService,
	}

	group.POST("/:organizationId/token", handler.createTokenHandler)
	group.GET("/:organizationId/token", handler.listTokensHandler)
	group.DELETE("/:organizationId/token/:id", handler.deleteTokenHandler)
}
