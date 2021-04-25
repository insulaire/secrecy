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
