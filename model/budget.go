package model

import "time"

type Budget struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Month         string    `json:"month" gorm:"not null"`
	Year          string    `json:"year" gorm:"not null"`
	TotalPrice    uint      `json:"total_price" gorm:"not null"`
	Food          uint      `json:"food" gorm:"not null"`
	Drink         uint      `json:"drink" gorm:"not null"`
	Book          uint      `json:"book" gorm:"not null"`
	Fashion       uint      `json:"fashion" gorm:"not null"`
	Furniture     uint      `json:"furniture" gorm:"not null"`
	GamesToys     uint      `json:"games_toys" gorm:"not null"`
	Beauty        uint      `json:"beauty" gorm:"not null"`
	EveryDayItems uint      `json:"every_dayItems" gorm:"not null"`
	Other         uint      `json:"other" gorm:"not null"`
	Notice        bool      `json:"notice" gorm:"not null"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	User          User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId        uint      `json:"user_id" gorm:"not null"`
}

type BudgetResponse struct {
	ID            uint      `json:"id"`
	Month         string    `json:"month"`
	Year          string    `json:"year"`
	TotalPrice    uint      `json:"total_price"`
	Food          uint      `json:"food"`
	Drink         uint      `json:"drink"`
	Book          uint      `json:"book"`
	Fashion       uint      `json:"fashion"`
	Furniture     uint      `json:"furniture"`
	GamesToys     uint      `json:"games_toys"`
	Beauty        uint      `json:"beauty"`
	EveryDayItems uint      `json:"every_dayItems"`
	Other         uint      `json:"other"`
	Notice        bool      `json:"notice"`
	CreatedAt     time.Time `json:"created_at"`
}
