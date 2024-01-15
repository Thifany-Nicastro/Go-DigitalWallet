package routes

import (
	"go-digitalwallet/controllers"
	"go-digitalwallet/infrastructure"
	"go-digitalwallet/repositories"
	"go-digitalwallet/services"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoutes(router *chi.Mux) {

	userRepository := &repositories.UserRepository{Dbinstance: infrastructure.DB}
	userService := &services.UserService{IUserRepository: userRepository}
	userController := controllers.UserController{IUserService: userService}

	router.HandleFunc("/users", userController.FindUser)
}
