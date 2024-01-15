package repositories

import (
	"go-digitalwallet/infrastructure"
	"go-digitalwallet/models"
)

type UserRepository struct {
	infrastructure.Dbinstance
}

func (repository *UserRepository) GetUserById(id string) (models.UserModel, error) {
	var user models.UserModel

	repository.Conn.First(&user, 1)

	return user, nil
}
