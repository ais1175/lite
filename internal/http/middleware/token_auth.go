package middleware

import (
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/fivemanage/lite/pkg/cache"
	"github.com/labstack/echo/v4"
)

func TokenAuth(tokenService *token.Service, cache *cache.Cache) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			// Check if the request has a valid token
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(401, echo.Map{
					"error": "Unauthorized: No token provided",
				})
			}

			token := tokenService.GetToken(ctx, token)

			// Proceed to the next handler
			return next(c)
		}
	}
}
