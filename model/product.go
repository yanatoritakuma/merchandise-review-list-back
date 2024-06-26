package model

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Stock       bool      `json:"stock" gorm:"not null"`
	Price       uint      `json:"price" gorm:"not null"`
	Review      float64   `json:"review" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Image       string    `json:"image" gorm:"not null"`
	Code        string    `json:"code" gorm:"not null"`
	Provider    string    `json:"provider" gorm:"not null"`
	TimeLimit   time.Time `json:"timeLimit" gorm:"not null"`
	User        User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId      uint      `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       bool      `json:"stock"`
	Price       uint      `json:"price"`
	Review      float64   `json:"review"`
	Url         string    `json:"url"`
	Image       string    `json:"image"`
	Provider    string    `json:"provider"`
	TimeLimit   time.Time `json:"timeLimit"`
	CreatedAt   time.Time
}

type ProductYearMonthResponse struct {
	TimeLimit time.Time `json:"timeLimit"`
}
