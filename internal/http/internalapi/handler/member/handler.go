package member

import (
	"github.com/fivemanage/lite/internal/http/appctx"
	"github.com/fivemanage/lite/internal/http/httputil"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (r *handler) listMembersHandler(c echo.Context) error {
	cc := c.(*appctx.Context)
	ctx := cc.Request().Context()

	id := cc.Param("id")

	members, err := r.memberService.ListMembers(ctx, id)
	if err != nil {
		logrus.WithError(err).Error("failed to list members")
		return echo.NewHTTPError(500, err)
	}

	return cc.JSON(200, httputil.Response(members))
}
