package utils

import (
"github.com/labstack/echo/v4"
"net/http"
)

type HttpResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func EchoHttpResponse(c echo.Context, code int, result interface{}) error {
	switch result.(type) {
	case HttpResponse:
		response := result.(HttpResponse)
		if len(response.Message) == 0 {
			if c == nil {
				return echo.NewHTTPError(code, HttpResponse{
					Message: http.StatusText(code),
					Code:    code,
				})
			}
			return c.JSON(code, HttpResponse{
				Message: http.StatusText(code),
				Code:    code,
			})
		}
		return c.JSON(code, HttpResponse{
			Message: response.Message,
			Code:    code,
		})
	}
	if c == nil {
		return echo.NewHTTPError(code, result)
	}
	return c.JSON(code, result)
}
