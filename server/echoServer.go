package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/yuta_2710/go-clean-arc-reviews/config"
	"github.com/yuta_2710/go-clean-arc-reviews/database"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/auth/handlers"
	AuthRouters "github.com/yuta_2710/go-clean-arc-reviews/modules/auth/routers"
	"github.com/yuta_2710/go-clean-arc-reviews/modules/auth/usecases"
	TokenRepo "github.com/yuta_2710/go-clean-arc-reviews/modules/token/repositories"
	UserHandler "github.com/yuta_2710/go-clean-arc-reviews/modules/users/handlers"
	UserRepository "github.com/yuta_2710/go-clean-arc-reviews/modules/users/repositories"
	UserRouters "github.com/yuta_2710/go-clean-arc-reviews/modules/users/routers"
	UserUsecase "github.com/yuta_2710/go-clean-arc-reviews/modules/users/usecases"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	root := s.initRootRoutes()
	// Health check
	s.app.GET("/health", func(c echo.Context) error {
		return c.String(200, "Health is OK")
	})

	// Initiliaze user https
	s.initUserHttps(root)
	s.initAuthHttps(root)

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (e *echoServer) initUserHttps(root *echo.Group) error {
	repo := UserRepository.NewUserPostgresRepository(e.db)
	usecase := UserUsecase.NewUserUsecaseImpl(repo)
	handler := UserHandler.NewUserHttp(usecase)
	UserRouters.InitUserRouters(handler, root)

	return nil
}

func (e *echoServer) initAuthHttps(root *echo.Group) error {
	userRepo := UserRepository.NewUserPostgresRepository(e.db)
	tokenRepo := TokenRepo.NewTokenPostgresRepository(e.db)
	usecase := usecases.NewAuthUsecaseImpl(userRepo, tokenRepo)
	handler := handlers.NewAuthHttp(usecase)
	AuthRouters.InitAuthRouters(handler, root)

	return nil
}

func (e *echoServer) initRootRoutes() *echo.Group {
	// if version == 1 {
	// 	v := os.Getenv("API_VERSION_1")
	// 	return e.Group("api/")
	// }
	return e.app.Group("api/v1")
}
