package repository

import (
	"fmt"
	"merchandise-review-list-backend/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IProductRepository interface {
	CreateProduct(product *model.Product) error
	UpdateTimeLimit(product *model.Product, userId uint, productId uint) error
	DeleteProduct(userId uint, productId uint) error
	GetMyProducts(product *[]model.Product, userId uint, page int, pageSize int) (int, error)
	GetMyProductsTimeLimitAll(product *[]model.Product, userId uint, page int, pageSize int, sort bool) (int, error)
	GetMyProductsTimeLimitYearMonth(product *[]model.Product, userId uint, yearMonth time.Time) error
	GetMyProductsTimeLimitDate(product *[]model.Product, userId uint, page int, pageSize int, date time.Time) (int, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) CreateProduct(product *model.Product) error {
	if err := pr.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) UpdateTimeLimit(product *model.Product, userId uint, productId uint) error {
	// 1日後に更新
	product.TimeLimit = product.TimeLimit.Add(24 * time.Hour)

	result := pr.db.Model(product).Clauses(clause.Returning{}).Where("id=? AND user_id=?", productId, userId).Updates(map[string]interface{}{
		"time_limit": product.TimeLimit,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (pr *productRepository) DeleteProduct(userId uint, productId uint) error {
	result := pr.db.Where("id=? AND user_id=?", productId, userId).Delete(&model.Product{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (pr *productRepository) GetMyProducts(product *[]model.Product, userId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := pr.db.Model(&model.Product{}).Where("user_id=?", userId).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := pr.db.Joins("User").Where("user_id=?", userId).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(product).Error; err != nil {
		return 0, err
	}
	return int(totalCount), nil
}

func (pr *productRepository) GetMyProductsTimeLimitAll(product *[]model.Product, userId uint, page int, pageSize int, sort bool) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	minimumTime := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

	if err := pr.db.Model(&model.Product{}).Where("user_id=? AND time_limit >= ?", userId, minimumTime).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	query := pr.db.Where("user_id=? AND time_limit >= ?", userId, minimumTime)

	if sort {
		query = query.Order("time_limit ASC")
	} else {
		query = query.Order("time_limit DESC")
	}

	if err := query.Offset(offset).Limit(pageSize).Find(product).Error; err != nil {
		return 0, err
	}

	return int(totalCount), nil
}

func (pr *productRepository) GetMyProductsTimeLimitYearMonth(product *[]model.Product, userId uint, yearMonth time.Time) error {

	startOfMonth := yearMonth.UTC()
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	// 年月が一致しているものを取得（同じ日にちは、統一している）
	if err := pr.db.Table("products").
		Select("DISTINCT ON (DATE_TRUNC('day', time_limit)) *").
		Where("user_id = ? AND time_limit >= ? AND time_limit < ?", userId, startOfMonth, endOfMonth).
		Order("DATE_TRUNC('day', time_limit), created_at DESC").
		Find(product).Error; err != nil {
		return err
	}

	return nil
}

func (pr *productRepository) GetMyProductsTimeLimitDate(product *[]model.Product, userId uint, page int, pageSize int, date time.Time) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	if err := pr.db.Model(&model.Product{}).Where("user_id=? AND DATE(time_limit)=?", userId, date).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := pr.db.Where("user_id=? AND DATE(time_limit)=?", userId, date).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(product).Error; err != nil {
		return 0, err
	}

	return int(totalCount), nil
}
