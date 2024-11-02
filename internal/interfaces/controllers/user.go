package controllers

import (
	"net/http"
	"simple-wallet/internal/application/usecases"
	"simple-wallet/internal/domain/models"
	"simple-wallet/internal/infrastructure/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService   *usecases.UserService
	walletService *usecases.WalletService
}

func NewUserController(userService *usecases.UserService, walletService *usecases.WalletService) *UserController {
	return &UserController{
		userService:   userService,
		walletService: walletService,
	}
}

func (c *UserController) Register(ctx echo.Context) error {
	user := models.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": "Invalid Request"})
	}

	if err := c.userService.HashPassword(user.Password, &user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Error hashing password"})
	}

	if err := c.userService.Register(&user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": err.Error()})
	}

	if err := c.walletService.Create(user.ID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"status": "success", "message": "Success to create user"})
}

func (c *UserController) Login(ctx echo.Context) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	loginRequest := LoginRequest{}
	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": "Invalid request"})
	}

	user, err := c.userService.GetByUsername(loginRequest.Username)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Invalid credentials"})
	}

	if err := c.userService.CheckPassword(loginRequest.Password, user); err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Invalid credentials"})
	}

	token, err := services.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Failed to generate token"})
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Login success",
		"data": map[string]string{
			"token": token,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}

func GetClaimsFromToken(c echo.Context) (*services.JWTClaims, error) {
	user := c.Get("user")
	if user == nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized access")
	}

	token := user.(*jwt.Token)
	claims, ok := token.Claims.(*services.JWTClaims)
	if !ok || !token.Valid {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	return claims, nil
}
