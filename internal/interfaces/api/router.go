package api

import (
	"simple-wallet/internal/application/usecases"
	"simple-wallet/internal/infrastructure/db"
	"simple-wallet/internal/infrastructure/repositories"
	"simple-wallet/internal/infrastructure/services"
	"simple-wallet/internal/interfaces/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func API(e *echo.Echo) {
	authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: services.JwtSecret,
		Claims:     &services.JWTClaims{},
	})

	// Connect to the database
	dbConn := db.Connect()

	userRepo := repositories.NewUserRepository(dbConn)
	walletRepo := repositories.NewWalletRepository(dbConn)

	userService := usecases.NewUserService(userRepo)
	walletService := usecases.NewWalletService(walletRepo, userRepo)

	userController := controllers.NewUserController(userService, walletService)
	walletController := controllers.NewWalletController(walletService)

	auth := e.Group("auth")
	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)

	wallet := e.Group("wallet")
	wallet.Use(authMiddleware)
	wallet.POST("/transfer", walletController.TransferMoney)
	wallet.GET("/balance", walletController.GetBalance)
}
