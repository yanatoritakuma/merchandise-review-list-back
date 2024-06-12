package validator

import (
	"merchandise-review-list-backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IBudgetValidator interface {
	BudgetValidator(budget model.Budget) error
}

type budgetValidator struct{}

func NewBudgetValidator() IBudgetValidator {
	return &budgetValidator{}
}

func (bv *budgetValidator) BudgetValidator(budget model.Budget) error {
	return validation.ValidateStruct(&budget,
		validation.Field(
			&budget.Month,
			validation.Required.Error("month is required"),
			validation.RuneLength(1, 3).Error("limites max 3 char"),
			validation.In("1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "all").Error("must be a valid month between 1 and 12 or all"),
		),
		validation.Field(
			&budget.Year,
			validation.Required.Error("month is required"),
			validation.RuneLength(1, 4).Error("limites max 4 char"),
		),
	)
}
