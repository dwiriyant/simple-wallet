package routes

import (
	"simple-wallet/controllers"
	"simple-wallet/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func API(e *echo.Echo) {
	authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: services.JwtSecret,
		Claims:     &services.JWTClaims{},
	})

	auth := e.Group("auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	wallet := e.Group("wallet")
	wallet.Use(authMiddleware)
	wallet.POST("/transfer", controllers.TransferMoney)
	wallet.GET("/balance", controllers.GetBalance)
}
