package http

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/service/authservice"
	"github.com/labstack/echo/v4"
)

func (r *Server) authRouterGroup(group *echo.Group, authService *authservice.Auth) {
	authGroup := group.Group("/auth")

	authGroup.POST("/register", func(c echo.Context) error {
		ctx := context.Background()

		var register api.RegisterRequest
		if err := BindAndValidate(c, &register); err != nil {
			return err
		}

		sessionID, err := authService.RegisterUser(ctx, &register)
		if err != nil {
			if errors.Is(err, authservice.ErrSessionExpired{}) {
			}
			log.Println("error: ", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to register user")
		}

		sessionCookie := authService.CreateSessionCookie(sessionID)
		c.SetCookie(sessionCookie)
		return c.Redirect(http.StatusSeeOther, "/app")
	})

	authGroup.GET("/session", func(c echo.Context) error {
		ctx := context.Background()

		sessionCookie, err := c.Cookie("fmlite_session")
		if err != nil {
			return c.JSON(http.StatusForbidden, "failed to find session cookie")
		}

		user, err := authService.UserBySession(ctx, sessionCookie.Value)
		if err != nil {
			if errors.Is(err, authservice.ErrSessionExpired{}) {
				return c.Redirect(http.StatusUnauthorized, "/auth")
			}
			return c.JSON(http.StatusForbidden, "failed to find session cookie")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"status": "ok",
			"data":   user,
		})
	})

	authGroup.POST("/login", func(c echo.Context) error {
		authService.LoginUser()

		return nil
	})
}
