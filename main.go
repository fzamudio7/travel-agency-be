package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const serverPort = "3333"

func getMessage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func main() {
	e := echo.New()
	e.GET("/message", getMessage)

	e.Logger.Fatal(e.Start(":" + serverPort))
}
