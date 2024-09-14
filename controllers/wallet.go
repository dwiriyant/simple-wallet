package controllers

import (
	"net/http"
	"simple-wallet/db"
	"simple-wallet/models"

	"github.com/labstack/echo/v4"
)

type transferRequest struct {
	Username string  `json:"username"`
	Amount   float64 `json:"amount"`
}

func TransferMoney(c echo.Context) error {
	userID, err := GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Unauthorized"})
	}

	transferRequest := transferRequest{}
	if err := c.Bind(&transferRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": "Invalid request"})
	}

	var fromWallet models.Wallet
	if err := db.DB.Where("user_id = ?", userID).First(&fromWallet).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"status": "error", "message": "Sender wallet not found"})
	}

	var toWallet models.Wallet
	if err := db.DB.Joins("User").Where("User.username = ?", transferRequest.Username).First(&toWallet).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"status": "error", "message": "Recipient wallet not found"})
	}

	if err := fromWallet.Transfer(db.DB, transferRequest.Amount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": err.Error()})
	}

	toWallet.Balance += transferRequest.Amount
	if err := db.DB.Save(&toWallet).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Failed to update recipient wallet"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success", "message": "Transfer Success"})
}
