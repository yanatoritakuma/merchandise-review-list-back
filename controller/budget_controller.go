package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IBudgetController interface {
	CreateBudget(c echo.Context) error
}

type budgetController struct {
	bu usecase.IBudgetUsecase
}

func NewBudgetController(bu usecase.IBudgetUsecase) IBudgetController {
	return &budgetController{bu}
}

func (bc *budgetController) CreateBudget(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	budget := model.Budget{}
	if err := c.Bind(&budget); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	budget.UserId = uint(userId.(float64))

	budgetRes, err := bc.bu.CreateProduct(budget)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, budgetRes)
}
