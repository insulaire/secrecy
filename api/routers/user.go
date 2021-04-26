package routers

import (
	"secrecy/api/handles"

	"github.com/gin-gonic/gin"
)

func InitUserRouters(e *gin.Engine) {
	g := e.Group("user")
	user := handles.NewUserController()
	g.POST("/login", user.Login())
	g.POST("/registered", user.Registered())
	g.POST("/list", user.QueryUserList())
}
