package models

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Any interface{}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewResponse(code int, message string, content Any) Response {
	if code >= 300 {
		err, ok := content.(*echo.HTTPError)

		if ok {
			return Response{
				Code:    code,
				Message: message,
				Error:   fmt.Sprintf("%v", err.Message),
			}
		}

		return Response{
			Code:    code,
			Message: message,
			Error:   fmt.Sprintf("%v", content),
		}
	}

	return Response{
		Code:    code,
		Message: message,
		Data:    content,
	}
}
