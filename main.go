package main

import (
	"github.com/labstack/echo/v4"
	"lbn/configs"
	"net/http"
)

func main() {
	e := echo.New()
	configs.ConnectDB()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
