package main

import (
	"github.com/gin-gonic/gin"
	"newsCenter/cmd/interfaces/user"
	"newsCenter/cmd/middlewares"
)

func initRouter(r *gin.Engine) {
	router := r.Group("/newsCenter")
	usr := router.Group("user")
	usr.POST("/register", user.Register)
	usr.POST("/login", user.Login)

	art := router.Group("/article")
	art.Use(middlewares.TokenVerify())
	//art.POST("/publish", article.Publish)
}
