package infrastructure

import (
	"golang-clean-architecture-example/api"
	"golang-clean-architecture-example/controllers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	userController *controllers.UserController
}

func NewServer(
	userController *controllers.UserController,
) *Server {
	return &Server{
		userController: userController,
	}
}

func (s *Server) UpdateUserName(ctx echo.Context, userID string) error {
	return s.userController.UpdateUserName(ctx, userID)
}

func InitRouter() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	container := BuildContainer()

	var server *Server
	if err := container.Invoke(func(s *Server) {
		server = s
	}); err != nil {
		log.Fatalf("Error resolving dependencies: %v", err)
	}

	api.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8080"))
}
