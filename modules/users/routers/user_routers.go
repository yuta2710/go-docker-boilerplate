package routers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/handlers"
)

func InitUserRouters(http handlers.UserHandler, routes *echo.Group) {
	routers := routes.Group("/users")
	fmt.Println("Registering POST /api/v1/users")
	routers.POST("/", http.CreateNewUser)
	routers.GET("/:id", http.GetUserById)
}
