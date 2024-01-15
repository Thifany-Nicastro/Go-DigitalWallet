package interfaces

import "go-digitalwallet/models"

type IUserService interface {
	GetUser(id string) (models.UserModel, error)
}
