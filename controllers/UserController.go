package controllers

import (
	"encoding/json"
	"go-digitalwallet/interfaces"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) FindUser(res http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "id")

	user, _ := controller.GetUser(userId)
	// if err != nil {
	// 	//Handle error
	// }

	json.NewEncoder(res).Encode(user)
}
