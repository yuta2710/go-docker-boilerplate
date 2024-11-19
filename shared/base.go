package shared

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type RouterGroup struct {
	Protected *echo.Group `json:"protected"`
	Public    *echo.Group `json:"public"`
}

func Response(c echo.Context, success bool, statusCode int, message string, data interface{}, headers map[string]string) error {
	if headers != nil {
		MapHeaders(c, headers)
	}

	if data != nil {
		return c.JSON(statusCode, BaseResponse{Success: success, Message: message, Status: statusCode, Data: data})
	}

	return c.JSON(statusCode, BaseResponse{Success: success, Message: message, Status: statusCode})
}

func MapHeaders(ctx echo.Context, headers map[string]string) error {
	for key, val := range headers {
		ctx.Response().Header().Set(key, val)
	}

	return nil
}
