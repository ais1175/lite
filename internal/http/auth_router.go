package http

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/labstack/echo/v4"
)

func (r *Server) authRouterGroup(group *echo.Group, authService *auth.Auth) {
	authGroup := group.Group("/auth")

	authGroup.POST("/register", func(c echo.Context) error {
		ctx := context.Background()

		var register api.RegisterRequest
		if err := BindAndValidate(c, &register); err != nil {
			return err
		}

		sessionID, err := authService.RegisterUser(ctx, &register)
		if err != nil {
			if errors.Is(err, auth.ErrSessionExpired{}) {
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
			if errors.Is(err, auth.ErrSessionExpired{}) {
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
		ctx := context.Background()

		var login api.LoginRequest
		if err := BindAndValidate(c, &login); err != nil {
			return err
		}

		sessionID, err := authService.LoginUser(ctx, login.Username, login.Password)
		if err != nil {
			// this might not be the way we want to handle these errors
			// might be better to send some sort of code instead that we can map on the client?
			// as this could get veryyy long some places
			// errors.As() is also an option
			if errors.Is(err, auth.ErrUserCredentials{}) {
				return c.JSON(http.StatusForbidden, echo.Map{
					"status":  "error",
					"message": "The username or password is wrong. Please try again",
				})
			}
			return c.JSON(http.StatusForbidden, echo.Map{
				"status":  "error",
				"message": err.Error(),
			})

		}

		sessionCookie := authService.CreateSessionCookie(sessionID)
		c.SetCookie(sessionCookie)
		return c.Redirect(http.StatusSeeOther, "/app")
	})
}
