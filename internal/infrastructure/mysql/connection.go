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
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
}

var once sync.Once

var MysqlConnectPool pool.Pool

func InitConnectPool(config *DBConfig) {
	once.Do(func() {
		MysqlConnectPool = pool.NewPool(pool.PoolConfig{
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

func GetDB() *gorm.DB {
	db, _ := MysqlConnectPool.Get().(*gorm.DB)
	return db
}

func newMysqlDB(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.String())
	return db, err
}
