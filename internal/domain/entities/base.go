package entities

import (
	"time"
)

type BaseEntity struct {
	Id       uint64 `gorm:"primary_key"`
	CreateAt time.Time
	CreateBy string
	DeleteAt time.Time
	DeleteBy string
}
