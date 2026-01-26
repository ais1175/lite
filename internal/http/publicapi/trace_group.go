package publicapi

import "github.com/labstack/echo/v4"

func registerTraceApi(group *echo.Group) {
	h := &traceHandler{}
	group.GET("/traces", h.getTraces)
}

type traceHandler struct{}

// getTraces godoc
// @Summary      Get traces
// @Description  Get tracing information
// @Tags         public
// @Produce      json
// @Success      200  {object}  echo.Map
// @Router       /traces [get]
func (h *traceHandler) getTraces(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}
