package main

import (
	"clean_sample/controller"
	"clean_sample/database"
	"clean_sample/repository"
	"clean_sample/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.InitDB()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
	r.POST("/users/bulk", userController.CreateUsersBulk)

	r.Run()
}
