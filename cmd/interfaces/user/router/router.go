package router

import (
	"github.com/gin-gonic/gin"
	"newsCenter/cmd/interfaces/user"
	"newsCenter/cmd/middlewares/tokenVerify"
)

func InitUserRouter(r *gin.Engine) {
	router := r.Group("newsCenter/user")
	router.POST("/register", user.Register)
	router.POST("/login", user.Login)
	router.Use(tokenVerify.TokenVerify())
	router.POST("/logout", user.LogoutJWT)
}
