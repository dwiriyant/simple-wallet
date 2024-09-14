package main

import (
	"net/http"
	"simple-wallet/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Simple Wallet REST API")
	})
	e.Logger.Fatal(e.Start(":8081"))
}
