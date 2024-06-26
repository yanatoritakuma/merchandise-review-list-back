package model

import "time"

type Like struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	ReviewPost ReviewPost `json:"reviewPost" gorm:"foreignKey:PostId; constraint:OnDelete:CASCADE"`
	PostId     uint       `json:"post_id" gorm:"not null"`
	User       User       `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId     uint       `json:"user_id" gorm:"not null"`
	PostUserId uint       `json:"post_user_id" gorm:"not null"`
}

type LikeResponse struct {
	ID     uint `json:"id"`
	UserId uint `json:"user_id"`
}
