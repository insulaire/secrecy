package depositaries

import (
	"secrecy/internal/domain/entities"
)

type IUser interface {
	//CheckPassword(name, password string) bool
	//GetInfos(user []entities.User)
	GetUserInfo(name string) (*entities.User, error)
}

// func NewUser() interfaces.IUser {
// 	return &User{}
// }
