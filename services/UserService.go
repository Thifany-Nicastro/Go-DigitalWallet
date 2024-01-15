package services

import (
	"go-digitalwallet/interfaces"
	"go-digitalwallet/models"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) GetUser(id string) (models.UserModel, error) {
	user, err := service.GetUserById(id)

	return user, err
}
