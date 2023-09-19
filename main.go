package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const serverPort = "3333"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":" + serverPort))
}
