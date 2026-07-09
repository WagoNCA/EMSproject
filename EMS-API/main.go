package main

import (
	"EMSproject/database"

	"github.com/labstack/echo/v5"
)

func main() {
	database.Connect()

	e := echo.New()

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error(err.Error())
	}
}
