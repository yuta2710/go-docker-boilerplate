package routers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/auth/handlers"
)

func InitAuthRouters(http handlers.AuthHandler, routes *echo.Group) {
	routers := routes.Group("/auth")
	fmt.Println("Login started....")

	routers.POST("/login", http.Login)
}
