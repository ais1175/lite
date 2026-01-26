package httputil

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type ResponseData struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorResponseData struct {
	Status string    `json:"status"`
	Data   ErrorData `json:"data"`
}

type ErrorData struct {
	Message string `json:"message"`
}

func Response(data any) *ResponseData {
	return &ResponseData{
		Status: "ok",
		Data:   data,
	}
}

func ErrorResponse(message string) *ErrorResponseData {
	return &ErrorResponseData{
		Status: "error",
		Data: ErrorData{
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
