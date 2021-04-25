package depositaries

import (
	"secrecy/internal/domain/entities"
)

type IUser interface {
	GetUserInfo(name string) (*entities.User, error)
}
