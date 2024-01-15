package main

import (
	"log"
	"net/http"

	"go-digitalwallet/infrastructure"
)

func main() {
	infrastructure.ConnectDb()

	log.Print("server has started")
	http.ListenAndServe(":3000", InitRouter())
}
