package member

import (
	"github.com/fivemanage/lite/internal/service/member"
	"github.com/labstack/echo/v4"
)

type handler struct {
	memberService *member.Service
}

func RegisterRoutes(group *echo.Group, memberService *member.Service) {
	h := handler{
		memberService: memberService,
	}

	group.GET("/organization/:id/member", h.listMembersHandler)
}
