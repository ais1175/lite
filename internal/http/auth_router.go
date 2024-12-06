package http

import (
	"context"
	"fmt"
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
			log.Println("error: ", err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to register user")
		}

		fmt.Println("session id", sessionID)

		sessionCookie := &http.Cookie{
			Name:     "fm_session",
			Value:    sessionID,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			HttpOnly: true,
			Path:     "/",
			MaxAge:   3600,
		}
		c.SetCookie(sessionCookie)
		return c.Redirect(http.StatusPermanentRedirect, "/app")
	})

	authGroup.GET("/session", func(c echo.Context) error {
		return nil
	})

	authGroup.POST("/login", func(c echo.Context) error {
		authService.LoginUser()

		return nil
	})

	authGroup.GET("/callback/github", func(c echo.Context) error {
		code := c.QueryParam("code")

		token := authService.Callback(code)

		return c.JSON(200, echo.Map{
			"token":          token.AccessToken,
			"resfresh_token": token.RefreshToken,
			"expiry":         token.Expiry,
		})
	})
}
