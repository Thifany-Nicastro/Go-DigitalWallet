package controllers

import (
	"encoding/json"
	"go-digitalwallet/interfaces"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) FindUser(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))

	user, err := controller.GetUser(id)
	if err != nil {
		res.WriteHeader(404)
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	json.NewEncoder(res).Encode(user)
}
