package user

import (
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	router := r.Group("newsCenter/user")
	router.POST("/register", Register)
	router.POST("/login", Login)
}
