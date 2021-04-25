package mysql

import (
	"github.com/jinzhu/gorm"
)

type Context struct {
}

// func (this *Context) handle(handle func(*gorm.DB)) {
// 	db, _ := mysqlConnectPool.Get().(*gorm.DB)
// 	defer mysqlConnectPool.Push(db)
// 	handle(db)
// }

func (this *Context) Handle(handle func(*gorm.DB) error) error {
	db, _ := mysqlConnectPool.Get().(*gorm.DB)
	defer mysqlConnectPool.Push(db)
	return handle(db)
}
