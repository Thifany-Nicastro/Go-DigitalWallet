package main

import (
	"go-digitalwallet/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	routes.SetupUserRoutes(router, InjectUserController())

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return router
}
