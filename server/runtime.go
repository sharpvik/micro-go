package server

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sharpvik/micro-go/database/names"
)

type runtime struct {
	names   names.Repo
	message string
}

func Runtime(db *sqlx.DB) *runtime {
	return &runtime{
		names:   names.NewRepo(db),
		message: "hello world",
	}
}

func (rt *runtime) Echo() (e *echo.Echo) {
	e = echo.New()
	e.Any("/", rt.greeting)
	e.POST("/:name", rt.remember)
	return
}

func (rt *runtime) greeting(c echo.Context) error {
	return c.String(http.StatusOK, rt.message)
}

func (rt *runtime) remember(c echo.Context) error {
	name := c.Param("name")
	if err := rt.names.Add(name); err != nil {
		return c.String(http.StatusInternalServerError, "failed to remember you")
	}
	return c.String(http.StatusOK, fmt.Sprintln("hello,", name))
}
