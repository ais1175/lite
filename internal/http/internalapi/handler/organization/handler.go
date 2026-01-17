package organization

import (
	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/http/appctx"
	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/fivemanage/lite/internal/http/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (r *handler) createOrganizationHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	var data api.CreateOrganizationRequest
	if err := validator.BindAndValidate(cc, &data); err != nil {
		logrus.WithError(err).Error("failed to bind and validate token")
		return echo.NewHTTPError(500, err)
	}

	user := cc.User()

	organization, err := r.organizationService.CreateOrganization(ctx, &data, user.ID)
	if err != nil {
		logrus.WithError(err).Error("failed to create organization")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(organization))
}

func (r *handler) listOrganizationsHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	organizations, err := r.organizationService.ListOrganizations(ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to list organizations")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(organizations))
}

func (r *handler) getOrganizationHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	id := cc.Param("id")

	organization, err := r.organizationService.FindOrganizationByID(ctx, id)
	if err != nil {
		logrus.WithError(err).Error("failed to find organization")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(organization))
}

func (r *handler) getOrganizationStatsHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	id := cc.Param("id")

	stats, err := r.organizationService.GetStats(ctx, id)
	if err != nil {
		logrus.WithError(err).Error("failed to get organization stats")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(stats))
}
