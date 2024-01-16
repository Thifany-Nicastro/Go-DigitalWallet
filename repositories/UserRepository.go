package repositories

import (
	"go-digitalwallet/infrastructure"
	"go-digitalwallet/models"
)

type UserRepository struct {
	infrastructure.Dbinstance
}

func (repository *UserRepository) GetUserById(id int) models.UserModel {
	var user models.UserModel

	repository.Conn.First(&user, id)

	return user
}
