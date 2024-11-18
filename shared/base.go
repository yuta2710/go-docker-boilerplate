package shared

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func Response(c echo.Context, success bool, statusCode int, message string, data interface{}) error {
	if data != nil {
		return c.JSON(statusCode, BaseResponse{Success: success, Message: message, Status: statusCode, Data: data})
	}

	return c.JSON(statusCode, BaseResponse{Success: success, Message: message, Status: statusCode})
}
