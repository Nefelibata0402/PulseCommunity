package main

import (
	"github.com/gin-gonic/gin"
	"newsCenter/cmd/interfaces/user"
	"newsCenter/cmd/middlewares"
)

func initRouter(r *gin.Engine) {
	router := r.Group("/newsCenter")
	router.POST("/user/register", user.Register)
	router.POST("/user/login", user.Login)
	router.Use(middlewares.TokenVerify())
}
