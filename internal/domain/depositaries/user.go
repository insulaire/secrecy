package depositaries

import (
	"secrecy/internal/domain/entities"
)

type IUser interface {
	GetUserInfo(string) (*entities.User, error)
	InsertUser(entities.User) error
	QueryUserList() []entities.User
}
