package entities

import (
	"errors"
)

type User struct {
	BaseEntity
	Name     string
	Password string
}

func (user *User) Valid() error {
	if user.Name == "" {
		return errors.New("用户名称必填")
	}
	return nil
}
