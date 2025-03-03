package internalapi

import (
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/labstack/echo/v4"
)

func Add(group *echo.Group, authservice *auth.Auth, tokenservice *token.Service) {
	registerAuthApi(group, authservice)
	registerTokensApi(group, tokenservice)
}
