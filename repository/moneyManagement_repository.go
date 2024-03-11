package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IMoneyManagementRepository interface {
	CreateMoneyManagement(moneyManagement *model.MoneyManagement) error
}

type moneyManagementRepository struct {
	db *gorm.DB
}

func NewMoneyManagementRepository(db *gorm.DB) IMoneyManagementRepository {
	return &moneyManagementRepository{db}
}

func (mr *moneyManagementRepository) CreateMoneyManagement(moneyManagement *model.MoneyManagement) error {
	if err := mr.db.Create(moneyManagement).Error; err != nil {
		return err
	}
	return nil
}
