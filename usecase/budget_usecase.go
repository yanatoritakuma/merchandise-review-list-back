package usecase

import (
	"errors"
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
)

type IBudgetUsecase interface {
	CreateProduct(budget model.Budget) (model.BudgetResponse, error)
}

type budgetUsecase struct {
	br repository.IBudgetRepository
	bv validator.IBudgetValidator
}

func NweBudgetUsecase(br repository.IBudgetRepository, bv validator.IBudgetValidator) IBudgetUsecase {
	return &budgetUsecase{br, bv}
}

func (bu *budgetUsecase) CreateProduct(budget model.Budget) (model.BudgetResponse, error) {
	existingBudget, err := bu.br.SameYearMonth(budget.UserId, budget.Year, budget.Month)

	if err != nil {
		return model.BudgetResponse{}, err
	}

	if existingBudget != nil {
		return model.BudgetResponse{}, errors.New("duplicate budget")
	}

	if err := bu.bv.BudgetValidator(budget); err != nil {
		return model.BudgetResponse{}, err
	}

	if err := bu.br.CreateBudget(&budget); err != nil {
		return model.BudgetResponse{}, err
	}

	resBudget := model.BudgetResponse{
		ID:            budget.ID,
		Month:         budget.Month,
		Year:          budget.Year,
		TotalPrice:    budget.TotalPrice,
		Food:          budget.Food,
		Drink:         budget.Drink,
		Book:          budget.Book,
		Fashion:       budget.Fashion,
		Furniture:     budget.Furniture,
		GamesToys:     budget.GamesToys,
		Beauty:        budget.Beauty,
		EveryDayItems: budget.EveryDayItems,
		Other:         budget.Other,
		Notice:        budget.Notice,
		CreatedAt:     budget.CreatedAt,
	}
	return resBudget, nil
}
