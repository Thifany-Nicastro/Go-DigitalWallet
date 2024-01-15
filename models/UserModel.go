package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name string `json:"name"`
}

type Tabler interface {
	TableName() string
}

func (UserModel) TableName() string {
	return "user"
}
