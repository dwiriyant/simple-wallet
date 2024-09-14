package controllers

import (
	"net/http"
	"simple-wallet/db"
	"simple-wallet/models"
	"simple-wallet/services"

	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": "Invalid Request"})
	}

	if err := user.HashPassword(user.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Error hashing password"})
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success", "message": "Success to create user"})
}

func Login(c echo.Context) error {
	loginRequest := loginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": "Invalid request"})
	}

	var user models.User
	if err := db.DB.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Invalid credentials"})
	}

	if err := user.CheckPassword(loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Invalid credentials"})
	}

	token, err := services.GenerateJWT(int(user.ID), user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Failed to generate token"})
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Login success",
		"data": map[string]string{
			"token": token,
		},
	}

	return c.JSON(http.StatusOK, response)
}