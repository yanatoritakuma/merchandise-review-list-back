package model

import "time"

type MoneyManagement struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"not null"`
	Category   string    `json:"category" gorm:"not null"`
	UnitPrice  uint      `json:"unit_price" gorm:"not null"`
	Quantity   uint      `json:"quantity" gorm:"not null"`
	TotalPrice uint      `json:"total_price" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId     uint      `json:"user_id" gorm:"not null"`
}

type MoneyManagementResponse struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"not null"`
	Category   string    `json:"category" gorm:"not null"`
	UnitPrice  uint      `json:"unit_price" gorm:"not null"`
	Quantity   uint      `json:"quantity" gorm:"not null"`
	TotalPrice uint      `json:"total_price" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
}
