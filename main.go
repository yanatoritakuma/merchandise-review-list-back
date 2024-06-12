package main

import (
	"merchandise-review-list-backend/controller"
	"merchandise-review-list-backend/db"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/router"
	"merchandise-review-list-backend/usecase"
	"merchandise-review-list-backend/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NweUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	productValidator := validator.NewProductValidator()
	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NweProductUsecase(productRepository, productValidator)
	productController := controller.NewProductController(productUsecase)

	likeRepositor := repository.NewLikeRepository(db)
	likeUsecase := usecase.NewLikeUsecase(likeRepositor)
	likeController := controller.NewLikeController(likeUsecase)

	reviewPostValidator := validator.NewReviewPostValidator()
	reviewPostRepository := repository.NewPostRepository(db)
	reviewPostUsecase := usecase.NewReviewPostUsecase(reviewPostRepository, reviewPostValidator, likeRepositor)
	reviewPostController := controller.NewReviewPostController(reviewPostUsecase)

	commentValidator := validator.NewCommentValidator()
	commentRepository := repository.NewCommentRepository(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepository, commentValidator, reviewPostRepository)
	commentController := controller.NewCommentController(commentUsecase)

	moneyManagementRepository := repository.NewMoneyManagementRepository(db)
	moneyManagementValidator := validator.NewMoneyManagementValidator()
	moneyManagementUsecase := usecase.NewMoneyManagementUsecase(moneyManagementRepository, moneyManagementValidator)
	moneyManagementController := controller.NewMoneyManagementController(moneyManagementUsecase)

	budgetRepository := repository.NewBudgetRepository(db)
	budgetValidator := validator.NewBudgetValidator()
	budgetUsecase := usecase.NweBudgetUsecase(budgetRepository, budgetValidator)
	budgetController := controller.NewBudgetController(budgetUsecase)

	e := router.NewRouter(userController, productController, reviewPostController, likeController, commentController, moneyManagementController, budgetController)
	e.Logger.Fatal(e.Start(":8080"))
}
