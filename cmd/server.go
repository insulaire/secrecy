package main

import (
	"secrecy/api/routers"
	"secrecy/internal/infrastructure/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql.InitConnectPool(&mysql.DBConfig{
		User:     "root",
		Password: "root",
		DBName:   "db2",
		Host:     "127.0.0.1",
		Port:     3306,
	})
	g := gin.New()
	routers.InitUserRouters(g)

	g.Run(":8899")
}
