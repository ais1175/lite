package middleware

import (
	"github.com/fivemanage/lite/internal/http/appctx"
	"github.com/labstack/echo/v4"
)

func AppContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(&appctx.Context{Context: c})
	}
}
