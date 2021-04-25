package application

import (
	"secrecy/internal/infrastructure/depositaries"
)

type userApp struct {
	us depositaries.UserRepository
}

func NewUserApp() UserAppInterface {
	return &userApp{
		us: depositaries.UserRepository{},
	}
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	CheckPassword(name, password string) bool
}

func (this *userApp) CheckPassword(name, password string) bool {
	user, err := this.us.GetUserInfo(name)
	if err != nil {
		return false
	}
	return user.Password == password
}
