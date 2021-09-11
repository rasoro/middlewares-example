package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message": "noice"}`)
	})

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message": "ok"}`)
	})

	e.Start(":8090")
}
