package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/labstack/echo/v4"
)

func Session(authService *auth.Auth) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			fmt.Println(c.Request().URL.Path)
			// ignore /auth endpoints
			if strings.HasPrefix(c.Request().URL.Path, "/api/auth") {
				return next(c)
			}

			sessionCookie, err := c.Cookie("fmlite_session")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "failed to find session cookie")
			}

			session, err := authService.UserBySession(ctx, sessionCookie.Value)
			if err != nil {
				return err
			}

			c.Set("user", session)
			return next(c)
		}
	}
}
