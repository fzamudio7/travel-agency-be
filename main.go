package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const serverPort = "3333"

type Airline struct {
	Name string `json:"name"`
}

func getMessage(c echo.Context) error {
	airline := &Airline{Name: "Delta"}
	return c.JSON(http.StatusOK, airline)
}

func main() {
	e := echo.New()
	e.GET("/message", getMessage)

	e.Logger.Fatal(e.Start(":" + serverPort))
}
