package services

import (
	"errors"
	"go-digitalwallet/interfaces"
	"go-digitalwallet/models"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) GetUser(id int) (models.UserModel, error) {
	user := service.GetUserById(id)

	if user == (models.UserModel{}) {
		return user, errors.New("user not found")
	}

	return user, nil
}
