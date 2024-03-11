package main

import (
	"fmt"
	"merchandise-review-list-backend/db"
	"merchandise-review-list-backend/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Product{}, &model.ReviewPost{}, &model.Like{}, &model.Comment{}, &model.MoneyManagement{})
}
