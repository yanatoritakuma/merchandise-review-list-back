package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
)

type IMoneyManagementUsecase interface {
	CreateMoneyManagement(moneyManagement model.MoneyManagement) (model.MoneyManagementResponse, error)
}

type moneyManagementUsecase struct {
	mr repository.IMoneyManagementRepository
}

func NewMoneyManagementUsecase(
	mr repository.IMoneyManagementRepository,
) IMoneyManagementUsecase {
	return &moneyManagementUsecase{mr}
}

func (mu *moneyManagementUsecase) CreateMoneyManagement(moneyManagement model.MoneyManagement) (model.MoneyManagementResponse, error) {
	// todo:バリデーション未実装

	if err := mu.mr.CreateMoneyManagement(&moneyManagement); err != nil {
		return model.MoneyManagementResponse{}, err
	}

	resMoneyManagement := model.MoneyManagementResponse{
		ID:         moneyManagement.ID,
		Title:      moneyManagement.Title,
		Category:   moneyManagement.Category,
		UnitPrice:  moneyManagement.UnitPrice,
		Quantity:   moneyManagement.Quantity,
		TotalPrice: moneyManagement.TotalPrice,
		CreatedAt:  moneyManagement.CreatedAt,
	}
	return resMoneyManagement, nil
}
