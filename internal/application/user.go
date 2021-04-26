package application

import (
	"secrecy/internal/domain/depositaries"
	"secrecy/internal/domain/entities"
	de "secrecy/internal/infrastructure/depositaries"
)

type userApp struct {
	us depositaries.IUser
}

func NewUserApp() UserAppInterface {
	return &userApp{
		us: de.NewUserRepository(),
	}
}

var _ UserAppInterface = &userApp{}

type RegisteredDto struct {
	User     string
	Password string
}

type UserAppInterface interface {
	CheckPassword(name, password string) bool
	Registered(dto RegisteredDto) bool

	QueryUserList() []entities.User
}

func (this *userApp) CheckPassword(name, password string) bool {
	user, err := this.us.GetUserInfo(name)
	if err != nil {
		return false
	}
	return user.Password == password
}

func (this *userApp) Registered(dto RegisteredDto) bool {
	if err := this.us.InsertUser(entities.User{
		Name:     dto.User,
		Password: dto.Password,
	}); err != nil {
		return false
	}
	return true
}

func (this *userApp) QueryUserList() []entities.User {
	return this.us.QueryUserList()
}
