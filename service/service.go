package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Service struct {
	names NamesRepo
}

func New(names NamesRepo) *Service {
	return &Service{
		names: names,
	}
}

func (rt *Service) Server() (e *echo.Echo) {
	e = echo.New()

	// Middleware.
	e.Use(middleware.Recover(), middleware.Logger())

	// Endpoints.
	e.Any("/ping", rt.pong)
	e.POST("/:name", rt.remember)

	return
}

func (rt *Service) pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (rt *Service) remember(c echo.Context) error {
	name := c.Param("name")
	if err := rt.names.Add(name); err != nil {
		return c.String(http.StatusInternalServerError, "failed to remember you")
	}
	return c.String(http.StatusOK, fmt.Sprintln("hello,", name))
}
