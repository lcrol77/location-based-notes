package main

import (
	"github.com/labstack/echo/v4"
	"lbn/routes"
	"net/http"
)

func main() {
	e := echo.New()
	routes.NotesRoute(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
