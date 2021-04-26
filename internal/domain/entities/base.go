package entities

import (
	"encoding/binary"
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	Id        uint64     `gorm:"column:id;primary_key"`
	CreateAt  time.Time  `gorm:"column:create_at"`
	CreateBy  string     `gorm:"column:create_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at"` //gorm 自动实现软删除字段
	DeleteBy  string     `gorm:"column:delete_by"`
}

func (base *BaseEntity) InsertDefault(createBy string) {
	uu := uuid.New()
	base.Id = binary.BigEndian.Uint64(uu[0:8])
	base.CreateAt = time.Now().UTC()
	base.CreateBy = createBy
}

func (base *BaseEntity) DeleteDefault(deleteBy string) {
	base.DeleteBy = deleteBy
}
