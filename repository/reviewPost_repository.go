package repository

import (
	"fmt"
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IReviewPostRepository interface {
	CreateReviewPost(reviewPost *model.ReviewPost) error
	UpdateReviewPost(reviewPost *model.ReviewPost, userId uint, postId uint) error
	DeleteReviewPost(userId uint, postId uint) error
	GetMyReviewPosts(reviewPost *[]model.ReviewPost, userId uint, page int, pageSize int) (int, error)
	GetReviewPostById(reviewPost *model.ReviewPost, postId uint) error
	GetUserById(id uint) (*model.User, error)
	GetReviewPostLists(reviewPost *[]model.ReviewPost, category string, page int, pageSize int) (int, error)
	GetLikesByPostId(likes *[]model.Like, postId uint) error
	GetCommentsByPostId(comments *[]model.Comment, postId uint) error
}

type reviewPostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IReviewPostRepository {
	return &reviewPostRepository{db}
}

func (rr *reviewPostRepository) CreateReviewPost(reviewPost *model.ReviewPost) error {
	if err := rr.db.Create(reviewPost).Error; err != nil {
		return err
	}
	return nil
}

func (rr *reviewPostRepository) UpdateReviewPost(reviewPost *model.ReviewPost, userId uint, postId uint) error {
	result := rr.db.Model(reviewPost).Clauses(clause.Returning{}).Where("id=? AND user_id=?", postId, userId).Updates(map[string]interface{}{
		"title":    reviewPost.Title,
		"text":     reviewPost.Text,
		"image":    reviewPost.Image,
		"review":   reviewPost.Review,
		"category": reviewPost.Category,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (rr *reviewPostRepository) GetMyReviewPosts(reviewPost *[]model.ReviewPost, userId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := rr.db.Model(&model.ReviewPost{}).Where("user_id=?", userId).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := rr.db.Joins("User").Where("user_id=?", userId).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(reviewPost).Error; err != nil {
		return 0, err
	}
	return int(totalCount), nil
}

func (rr *reviewPostRepository) GetReviewPostById(reviewPost *model.ReviewPost, postId uint) error {
	if err := rr.db.First(reviewPost, postId).Error; err != nil {
		return err
	}
	return nil
}

func (rr *reviewPostRepository) GetUserById(id uint) (*model.User, error) {
	user := &model.User{}
	result := rr.db.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (rr *reviewPostRepository) DeleteReviewPost(userId uint, postId uint) error {
	result := rr.db.Where("id=? AND user_id=?", postId, userId).Delete(&model.ReviewPost{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (rr *reviewPostRepository) GetReviewPostLists(reviewPost *[]model.ReviewPost, category string, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	// categoryがallの場合は条件を無視して全てのレコードをカウント
	if category == "all" {
		if err := rr.db.Model(&model.ReviewPost{}).Count(&totalCount).Error; err != nil {
			return 0, err
		}
	} else {
		// 部分一致でカテゴリーを検索し、結果の数をカウント
		if err := rr.db.Model(&model.ReviewPost{}).Where("category LIKE ?", "%"+category+"%").Count(&totalCount).Error; err != nil {
			return 0, err
		}
	}

	// categoryがallの場合は条件を無視して全てのレコードを取得
	if category == "all" {
		if err := rr.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(reviewPost).Error; err != nil {
			return 0, err
		}
	} else {
		// 部分一致でカテゴリーを検索し、指定されたページとページサイズで結果を取得
		if err := rr.db.Where("category LIKE ?", "%"+category+"%").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(reviewPost).Error; err != nil {
			return 0, err
		}
	}

	return int(totalCount), nil
}

func (rr *reviewPostRepository) GetLikesByPostId(likes *[]model.Like, postId uint) error {
	return rr.db.Where("post_id=?", postId).Find(likes).Error
}

func (rr *reviewPostRepository) GetCommentsByPostId(comments *[]model.Comment, postId uint) error {
	return rr.db.Where("post_id=?", postId).Find(comments).Error
}
