package main

import (
	"simple-wallet/db"
	"simple-wallet/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()

	e := echo.New()

	routes.API(e)

	e.Logger.Fatal(e.Start(":8081"))
}
