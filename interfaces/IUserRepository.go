package interfaces

import "go-digitalwallet/models"

type IUserRepository interface {
	GetUserById(id string) (models.UserModel, error)
}
