package routes

import (
	"go-digitalwallet/controllers"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoutes(router *chi.Mux, userController controllers.UserController) {
	router.HandleFunc("/users", userController.FindUser)
}
