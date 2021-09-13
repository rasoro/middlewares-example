package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func applicationJSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response()
		w.Header().Set("Content-Type", "application/json")
		return next(c)
	}
}

func basicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		if r.URL.Path == "/healthcheck" {
			next(c)
		}
		user, pass, ok := r.BasicAuth()
		if !ok || user != "admin" || pass != "admin" {
			return echo.ErrUnauthorized
		}
		c.Response().Header().Set(echo.HeaderWWWAuthenticate, "Basic realm=Restricted")
		return next(c)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(applicationJSON)
	e.Use(basicAuth)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message": "noice"}`)
	})

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message": "ok"}`)
	})

	e.Start(":8090")
}
