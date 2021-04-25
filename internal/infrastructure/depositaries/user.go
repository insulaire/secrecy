package depositaries

import (
	"log"
	"secrecy/internal/domain/depositaries"
	"secrecy/internal/domain/entities"
	"secrecy/internal/infrastructure/mysql"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

var _ depositaries.IUser = &UserRepository{}

func (this *UserRepository) GetUserInfo(name string) (*entities.User, error) {

	db := mysql.GetDB()
	db.AutoMigrate(&entities.User{})
	var user entities.User
	err := db.Where("name = ?", name).Find(&user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
