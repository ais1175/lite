package middleware

import (
	"fmt"
	"time"

	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/fivemanage/lite/pkg/cache"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func TokenAuth(tokenService *token.Service, memcache *cache.Cache) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			ctx := c.Request().Context()

			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(401, echo.Map{
					"error": "Unauthorized: No token provided",
				})
			}

			var tokenData *database.Token

			// this all feels kinda weird to me
			key := fmt.Sprintf(cache.TokenCacheKey, token)
			tokenCacheData, found := memcache.Get(key)
			if !found {
				tokenData, err = tokenService.GetToken(ctx, token)
				if err != nil {
					otelzap.L().Error("failed to ge token from service", zap.Error(err))
					return nil
				}

				memcache.Set(key, tokenData, 5*time.Minute)

				c.Set("org_id", tokenData.OrganizationID)
				return next(c)
			}

			tokenData, ok := tokenCacheData.(*database.Token)
			if !ok {
				fmt.Println("failed to assert token data")
				return nil
			}

			c.Set("org_id", tokenData.OrganizationID)
			return next(c)
		}
	}
}
