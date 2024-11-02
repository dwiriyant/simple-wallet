package controllers

import (
	"net/http"
	"simple-wallet/internal/application/usecases"

	"github.com/labstack/echo/v4"
)

type WalletController struct {
	walletService *usecases.WalletService
}

func NewWalletController(walletService *usecases.WalletService) *WalletController {
	return &WalletController{walletService: walletService}
}

var GetClaimsFunc = GetClaimsFromToken

func (c *WalletController) TransferMoney(ctx echo.Context) error {
	type TransferRequest struct {
		Username string `json:"username" validate:"required"`
		Amount   int    `json:"amount" validate:"required,numeric,gt=0"`
	}

	claims, err := GetClaimsFunc(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Unauthorized"})
	}

	transferRequest := TransferRequest{}
	if err := ctx.Bind(&transferRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": "Invalid request"})
	}

	if claims.Username == transferRequest.Username {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": "Cannot transfer wallet to own user"})
	}

	amount64 := float64(transferRequest.Amount)
	if err := c.walletService.Transfer(claims.ID, transferRequest.Username, amount64); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"status": "success", "message": "Transfer Success"})
}

func (c *WalletController) GetBalance(ctx echo.Context) error {
	claims, err := GetClaimsFunc(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"status": "error", "message": "Unauthorized"})
	}

	balance, err := c.walletService.GetBalance(claims.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": err.Error()})
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Success get balance",
		"data": map[string]float64{
			"balance": balance,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}
