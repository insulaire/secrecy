package handles

import (
	"net/http"
	"secrecy/internal/application"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	user application.UserAppInterface
}

func NewUserController() UserController {
	return UserController{
		user: application.NewUserApp(),
	}
}

type LoginRequest struct {
	User     string
	Password string
}

func (this *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &LoginRequest{}
		c.BindJSON(&req)
		result := this.user.CheckPassword(req.User, req.Password)
		c.JSON(http.StatusOK, result)
	}
}

type RegisteredRequest struct {
	User     string
	Password string
}

func (this *UserController) Registered() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &RegisteredRequest{}
		c.BindJSON(&req)
		result := this.user.Registered(application.RegisteredDto{
			User:     req.User,
			Password: req.Password,
		})
		c.JSON(http.StatusOK, result)
	}
}

func (this *UserController) QueryUserList() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := this.user.QueryUserList()

		c.JSON(http.StatusOK, result)
	}
}
