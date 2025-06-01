package middleware

import (
	"fmt"

	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/fivemanage/lite/pkg/cache"
	"github.com/labstack/echo/v4"
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

			key := fmt.Sprintf(cache.TokenCacheKey, token)
			tokenCacheData, found := memcache.Get(key)
			if !found {
				tokenData, err = tokenService.GetToken(ctx, token)
				if err != nil {
					return nil
				}

				memcache.Set(key, tokenData, 100)
			}

			tokenData, ok := tokenCacheData.(*database.Token)
			if !ok {
				return nil
			}

			c.Set("org_id", tokenData.OrganizationID)
			return next(c)
		}
	}
}
