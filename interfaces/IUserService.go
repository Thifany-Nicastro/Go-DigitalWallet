package interfaces

import "go-digitalwallet/models"

type IUserService interface {
	GetUser(id int) (models.UserModel, error)
}
