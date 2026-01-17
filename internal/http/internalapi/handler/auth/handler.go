package auth

import (
	"errors"
	"net/http"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/http/appctx"
	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/http/validator"
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/labstack/echo/v4"

	internalauth "github.com/fivemanage/lite/internal/auth"
)

func (r *handler) getSessionHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	sessionCookie, err := cc.Cookie(internalauth.SessionCookieName)
	if err != nil {
		return cc.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Already logged out",
		})
	}

	user, err := r.authService.UserBySession(ctx, sessionCookie.Value)
	if err != nil {
		return cc.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid session",
		})
	}

	return cc.JSON(http.StatusOK, httputil.Response(user))
}

func (r *handler) loginHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	var login api.LoginRequest
	if err := validator.BindAndValidate(cc, &login); err != nil {
		return err
	}

	sessionID, err := r.authService.LoginUser(ctx, login.Username, login.Password)
	if err != nil {
		// this might not be the way we want to handle these errors
		// might be better to send some sort of code instead that we can map on the client?
		// as this could get veryyy long some places
		// errors.As() is also an option
		if errors.Is(err, auth.ErrUserCredentials{}) {
			return c.JSON(http.StatusForbidden, echo.Map{
				"status": "error",
				// might need some constant shit or something idk
				"message": "The username or password is wrong. Please try again",
			})
		}
		// this also bad
		return c.JSON(http.StatusForbidden, echo.Map{
			"status":  "error",
			"message": err.Error(),
		})

	}

	sessionCookie := r.authService.CreateSessionCookie(sessionID)
	c.SetCookie(sessionCookie)
	return c.Redirect(http.StatusSeeOther, "/app")
}

func (r *handler) logoutHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	sessionCookie, err := cc.Cookie(internalauth.SessionCookieName)
	if err != nil {
		// better error handling here
		return cc.JSON(http.StatusOK, echo.Map{
			"message": "Already logged out",
		})
	}

	err = r.authService.LogoutUser(ctx, sessionCookie.Value)
	if err != nil {
		// here too
		return cc.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to logout",
		})
	}

	logoutCookie := r.authService.CreateLogoutCookie()
	cc.SetCookie(logoutCookie)

	// change this dumb ahhh response
	return cc.JSON(http.StatusOK, echo.Map{
		"message": "Logout successful",
	})
}
