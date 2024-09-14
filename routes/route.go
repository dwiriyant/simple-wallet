package routes

import (
	"simple-wallet/controllers"

	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {
	auth := e.Group("auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
}
