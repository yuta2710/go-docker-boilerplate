package routers

import (
	"github.com/labstack/echo/v4"
	custom_middleware "github.com/yuta_2710/go-clean-arc-reviews/middleware"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/auth/handlers"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/users/repositories"
)

func InitAuthRouters(http handlers.AuthHandler, userRepo repositories.UserRepository, routes *echo.Group) {
	// ROUTE_VAL := "/auth"
	publicRoutes := routes.Group("/auth")
	publicRoutes.POST("/login", http.Login)
	// fmt.Println("Login started....")

	protectedRoutes := routes.Group("/auth")
	protectedRoutes.Use(custom_middleware.Protect(userRepo))

	protectedRoutes.GET("/profile", http.Profile)
}
