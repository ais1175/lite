package http

import (
	"embed"
	nethttp "net/http"
	"strings"

	"github.com/fivemanage/lite/internal/http/internalapi"
	"github.com/fivemanage/lite/internal/http/publicapi"
	_validator "github.com/fivemanage/lite/internal/http/validator"
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
)

// This embed shit shouldn't be this far away.
// When we are ready, we can either build the React code directly into the server or root folder
// or just copy it in the Dockerfile.

// go:embed ../../../web/dist
var webContent embed.FS

type Server struct {
	Engine *echo.Echo
}

// TODO: Add sentry for monitoring. There should be an opt-out option.
func NewServer(
	authservice *auth.Auth,
	tokenservice *token.Service,
	fileservice *file.Service,
) *echo.Echo {
	app := echo.New()

	// not good, not bad
	app.Validator = &_validator.CustomValidator{Validator: validator.New()}
	app.Use(middleware.Recover())

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "web/dist",
		Filesystem: nethttp.FS(webContent),
		HTML5:      true,
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
	}))

	apiGroup := app.Group("/api")

	internalapi.Add(
		apiGroup,
		authservice,
		tokenservice,
	)
	publicapi.Add(apiGroup, fileservice)

	// app.imageRouterGroup(apiGroup)

	return app
}
