package main

import (
	"go-digitalwallet/controllers"
	"go-digitalwallet/infrastructure"
	"go-digitalwallet/repositories"
	"go-digitalwallet/services"
)

func InjectUserController() controllers.UserController {
	userRepository := &repositories.UserRepository{Dbinstance: infrastructure.DB}
	userService := &services.UserService{IUserRepository: userRepository}
	userController := controllers.UserController{IUserService: userService}

	return userController
}
