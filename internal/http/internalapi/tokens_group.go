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
		// this might not work well for telemetry
		ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
		defer cancel()

		var data api.CreateTokenRequest
		if err := validator.BindAndValidate(c, &data); err != nil {
			logrus.WithError(err).Error("failed to bind and validate token")
			return echo.NewHTTPError(500, err)
		}

		apiToken, err := tokenService.CreateToken(ctx, &data)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}

		// we return the apiToken as a one-time thing
		tokenResponse := &api.CreateTokenResponse{
			Token: apiToken,
		}

		return c.JSON(200, httputil.Response(tokenResponse))
	})

	group.GET("/token", func(c echo.Context) error {
		ctx := c.Request().Context()

		tokens, err := tokenService.ListTokens(ctx)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}

		return c.JSON(200, httputil.Response(tokens))
	})

	group.DELETE("/token/:id", func(c echo.Context) error {
		var err error
		ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
		defer cancel()

		tokenID := c.Param("id")

		err = tokenService.DeleteToken(ctx, tokenID)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}

		return c.JSON(200, httputil.Response(nil))
	})
}
