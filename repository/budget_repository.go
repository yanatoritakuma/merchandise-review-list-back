package repository

import (
	"errors"
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IBudgetRepository interface {
	CreateBudget(budget *model.Budget) error
	SameYearMonth(userId uint, year string, month string) (*model.Budget, error) //既に設定年月が存在しているか
}

type budgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) IBudgetRepository {
	return &budgetRepository{db}
}

func (br *budgetRepository) CreateBudget(budget *model.Budget) error {
	if err := br.db.Create(budget).Error; err != nil {
		return err
	}
	return nil
}

func (br *budgetRepository) SameYearMonth(userId uint, year string, month string) (*model.Budget, error) {
	budget := &model.Budget{}

	if err := br.db.Where("user_id=? AND year=? AND month=?", userId, year, month).First(budget).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 登録済みの設定金額が存在しない場合はnilを返す
			return nil, nil
		}
		return nil, err
	}

	return budget, nil
}
