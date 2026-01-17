package httputil

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type errorResponse struct {
	Status string    `json:"status"`
	Data   errorData `json:"data"`
}

type errorData struct {
	Message string `json:"message"`
}

func Response(data any) *response {
	return &response{
		Status: "ok",
		Data:   data,
	}
}

func ErrorResponse(message string) *errorResponse {
	return &errorResponse{
		Status: "error",
		Data: errorData{
			Message: message,
		},
	}
}

func QueryInt(c echo.Context, key string, defaultValue int) (int, error) {
	val := c.QueryParam(key)
	if val == "" {
		return defaultValue, nil
	}

	res, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue, err
	}

	return res, nil
}
