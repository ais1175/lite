package internalapi

import (
	"errors"
	"net/http"

	"github.com/fivemanage/lite/api"
	internalauth "github.com/fivemanage/lite/internal/auth"
	"github.com/fivemanage/lite/internal/http/validator"
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/labstack/echo/v4"
)

func registerAuthApi(group *echo.Group, authservice *auth.Auth) {
	group.GET("/auth/session", func(c echo.Context) error {
		user := c.Get(internalauth.UserContextKey)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"status": "ok",
			"data":   user,
		})
	})

	group.POST("/auth/login", func(c echo.Context) error {
		ctx := c.Request().Context()

		var login api.LoginRequest
		if err := validator.BindAndValidate(c, &login); err != nil {
			return err
		}

		sessionID, err := authservice.LoginUser(ctx, login.Username, login.Password)
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

		sessionCookie := authservice.CreateSessionCookie(sessionID)
		c.SetCookie(sessionCookie)
		return c.Redirect(http.StatusSeeOther, "/app")
	})

	group.POST("/auth/logout", func(c echo.Context) error {
		ctx := c.Request().Context()

		sessionCookie, err := c.Cookie(internalauth.SessionCookieName)
		if err != nil {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "Already logged out",
			})
		}

		err = authservice.LogoutUser(ctx, sessionCookie.Value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Failed to logout",
			})
		}

		logoutCookie := authservice.CreateLogoutCookie()
		c.SetCookie(logoutCookie)

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Logout successful",
		})
	})
}
