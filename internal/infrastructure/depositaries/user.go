package depositaries

import (
	"secrecy/internal/domain/depositaries"
	"secrecy/internal/domain/entities"
	"secrecy/internal/infrastructure/mysql"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	mysql.Context
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

var _ depositaries.IUser = &UserRepository{}

func init() {
	db := NewUserRepository()

	if err := db.Handle(func(db *gorm.DB) error {
		return db.AutoMigrate(&entities.User{}).Error
	}); err != nil {
		panic(err)
	}
}

func (this *UserRepository) GetUserInfo(name string) (*entities.User, error) {
	var user entities.User
	err := this.Handle(func(db *gorm.DB) error {
		return db.Where("name = ?", name).Find(&user).Error
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (this *UserRepository) InsertUser(user entities.User) error {
	if err := user.Valid(); err != nil {
		return err
	}
	user.InsertDefault(user.CreateBy)

	return this.Handle(func(db *gorm.DB) error {
		return db.Create(user).Error
	})
}

func (this *UserRepository) QueryUserList() []entities.User {
	result := []entities.User{}
	this.Handle(func(db *gorm.DB) error {
		db.Unscoped().Find(&result)
		return nil
	})
	return result
}
