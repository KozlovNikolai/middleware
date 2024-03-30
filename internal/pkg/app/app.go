package app

import (
	"fmt"
	"log"

	"github.com/KozlovNikolai/middleware/internal/endpoint"
	"github.com/KozlovNikolai/middleware/internal/mw"
	"github.com/KozlovNikolai/middleware/internal/service"
	"github.com/labstack/echo/v4"
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
	a.echo.Use(mw.RoleCheck)
	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server running. ")
	return nil
}
