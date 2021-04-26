package mysql

import (
	"fmt"
	"sync"

	"secrecy/pkg/pool"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	options  string
}

func (conf *DBConfig) String() string {
	//?parseTime=true 这段缺失会导致 gorm 解析不到 time.Time类型
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
}

var once sync.Once

var mysqlConnectPool pool.Pool

func InitConnectPool(config *DBConfig) {
	once.Do(func() {
		mysqlConnectPool = pool.NewPool(pool.PoolConfig{
			InitFn: func() interface{} {
				db, err := newMysqlDB(config)
				if err != nil {
					panic(err)
				}
				return db
			},
			Min: 5,
			Max: 10,
		})
	})
}

func newMysqlDB(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.String())
	return db, err
}
