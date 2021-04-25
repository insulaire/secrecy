package entities

import (
	"log"
	"secrecy/internal/infrastructure/mysql"
)

type User struct {
	BaseEntity
	Name     string
	Password string
}

func (user *User) CheckPassword(name, password string) bool {
	db := mysql.GetDB()
	db.AutoMigrate(&User{})
	var count int
	err := db.Where("name = ?", name).Where("password = ?", password).Find(&User{}).Count(&count).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}
