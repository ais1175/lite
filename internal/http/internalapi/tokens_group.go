package internalapi

import (
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/labstack/echo/v4"
)

func registerTokensApi(group *echo.Group, tokenservice *token.Service) {
	group.POST("/tokens", func(c echo.Context) error {
		return nil
	})
}
