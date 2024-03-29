package repository

import (
	"fmt"
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment *model.Comment) error
	DeleteComment(userId uint, id uint) error
	GetCommentsByPostId(comments *[]model.Comment, postId uint, page int, pageSize int) (int, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) CreateComment(comment *model.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) DeleteComment(userId uint, id uint) error {
	result := cr.db.Where("user_id=? AND id=?", userId, id).Delete(&model.Comment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (cr *commentRepository) GetCommentsByPostId(comments *[]model.Comment, postId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := cr.db.Where("post_id=?", postId).Model(&model.Comment{}).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := cr.db.Where("post_id=?", postId).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(comments).Error; err != nil {
		return 0, err
	}

	return int(totalCount), nil
}
