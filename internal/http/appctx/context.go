package appctx

import (
	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/auth"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func (c *Context) User() *api.User {
	u, _ := c.Get(auth.UserContextKey).(*api.User)
	return u
}

func (c *Context) OrganizationID() string {
	orgID, _ := c.Get(auth.OrgIDContextKey).(string)
	return orgID
}
