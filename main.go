package main

import (
	"log"
	"net/http"

	"go-digitalwallet/infrastructure"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	infrastructure.ConnectDb()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Print("server has started")
	http.ListenAndServe(":3000", r)
}
