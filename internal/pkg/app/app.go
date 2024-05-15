package app

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/unreal-kz/umbrella-test-task/internal/app/endpoint"
	"github.com/unreal-kz/umbrella-test-task/internal/app/service"
	"github.com/unreal-kz/umbrella-test-task/internal/middleware"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()
	a.echo.Use(middleware.RoleCheck)
	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running...")
	err := a.echo.Start(":8080")
	if err != nil {
		a.echo.Logger.Fatal(err)
	}
	return nil
}
