package http

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/fivemanage/lite/internal/http/internalapi"
	"github.com/fivemanage/lite/internal/http/publicapi"
	_validator "github.com/fivemanage/lite/internal/http/validator"
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/fivemanage/lite/internal/service/organization"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/fivemanage/lite/pkg/cache"
	"github.com/labstack/echo/v4"
	"github.com/spf13/afero/mem"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed dist/*
var webContent embed.FS

type Server struct {
	Engine *echo.Echo
}

// TODO: Add sentry for monitoring. There should be an opt-out option.
func NewServer(
	authService *auth.Auth,
	tokenService *token.Service,
	fileService *file.Service,
	organizationService *organization.Service,
	memcache *cache.Cache,
) *echo.Echo {
	app := echo.New()
	app.Debug = true

	// not good, not bad
	app.Validator = &_validator.CustomValidator{Validator: validator.New()}
	app.Use(middleware.Recover())

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: getFileSystem("dist"),
		HTML5:      true,
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
	}))

	apiGroup := app.Group("/api")
	dashApi := apiGroup.Group("/dash")

	internalapi.Add(
		dashApi,
		authService,
		tokenService,
		organizationService,
		fileService,
	)
	publicapi.Add(apiGroup, fileService, tokenService, memcache)

	return app
}

func getFileSystem(path string) http.FileSystem {
	fs, err := fs.Sub(webContent, path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Serving static files from", path)

	return http.FS(fs)
}
