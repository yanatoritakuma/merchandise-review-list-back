package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IMoneyManagementController interface {
	CreateMoneyManagement(c echo.Context) error
}

type moneyManagementController struct {
	mu usecase.IMoneyManagementUsecase
}

func NewMoneyManagementController(mu usecase.IMoneyManagementUsecase) IMoneyManagementController {
	return &moneyManagementController{mu}
}

func (mc *moneyManagementController) CreateMoneyManagement(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	moneyManagement := model.MoneyManagement{}
	if err := c.Bind(&moneyManagement); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	moneyManagement.UserId = uint(userId.(float64))
	moneyManagementRes, err := mc.mu.CreateMoneyManagement(moneyManagement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, moneyManagementRes)
}
