package interfaces

import "go-digitalwallet/models"

type IUserRepository interface {
	GetUserById(id int) models.UserModel
}
