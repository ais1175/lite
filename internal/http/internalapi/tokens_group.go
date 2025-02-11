package internalapi

import (
	"context"
	"time"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/http/validator"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func registerTokensApi(group *echo.Group, tokenService *token.Service) {
	group.POST("/token", func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
		defer cancel()

		var data api.CreateTokenRequest
		if err := validator.BindAndValidate(c, &data); err != nil {
			logrus.WithError(err).Error("failed to bind and validate token")
			return c.JSON(500, err)
		}

		apiToken, err := tokenService.CreateToken(ctx, &data)
		if err != nil {
			return c.JSON(500, nil)
		}

		tokenResponse := &api.CreateTokenResponse{
			Token: apiToken,
		}

		return c.JSON(200, httputil.Response(tokenResponse))
	})
}
