package entities

import (
	"errors"
)

type User struct {
	BaseEntity
	Name     string `gorm:"column:name;unique_index"`
	Password string `gorm:"column:password"`
}

func (user *User) Valid() error {
	if user.Name == "" {
		return errors.New("用户名称必填")
	}
	return nil
}
