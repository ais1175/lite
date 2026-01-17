package token

import (
	"context"
	"time"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/http/appctx"
	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/http/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (r *handler) createTokenHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx, cancel := context.WithTimeout(cc.Request().Context(), 5*time.Second)
	defer cancel()

	organizationID := cc.OrganizationID()

	var data api.CreateTokenRequest
	if err := validator.BindAndValidate(cc, &data); err != nil {
		logrus.WithError(err).Error("failed to bind and validate token")
		return echo.NewHTTPError(500, err)
	}

	data.OrganizationID = organizationID
	apiToken, err := r.tokenService.CreateToken(ctx, &data)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}

	// we return the apiToken as a one-time thing
	tokenResponse := &api.CreateTokenResponse{
		Token: apiToken,
	}

	return cc.JSON(200, httputil.Response(tokenResponse))
}

func (r *handler) listTokensHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	organizationID := cc.OrganizationID()

	tokens, err := r.tokenService.ListTokens(ctx, organizationID)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(tokens))
}

func (r *handler) deleteTokenHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	tokenID := cc.Param("id")

	err := r.tokenService.DeleteToken(ctx, tokenID)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(nil))
}
