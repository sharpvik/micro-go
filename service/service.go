package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Service struct {
	users Users
}

func New(names Users) *Service {
	return &Service{
		users: names,
	}
}

func (s *Service) Server() (e *echo.Echo) {
	e = echo.New()

	// Middleware.
	e.Use(middleware.Recover(), middleware.Logger())

	// Endpoints.
	e.Any("/ping", s.pong)
	e.POST("/signup", s.remember)

	return
}

func (s *Service) pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (s *Service) remember(c echo.Context) (err error) {
	var body SignUpRequest
	if err = json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return c.String(http.StatusBadRequest,
			fmt.Sprintln("failed to decode request:", err))
	}
	if err := s.users.Add(body.Username, body.Password); err != nil {
		return c.String(http.StatusInternalServerError, "failed to sign you up")
	}
	return c.String(http.StatusOK, fmt.Sprintln("hello,", body.Username))
}
