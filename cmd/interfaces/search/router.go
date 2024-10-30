package search

import (
	"github.com/gin-gonic/gin"
	"pulseCommunity/cmd/middlewares/tokenVerify"
)

func InitSearchRouter(r *gin.Engine) {
	router := r.Group("newsCenter/search")
	router.Use(tokenVerify.TokenVerify())
	router.POST("/search", Search)
}
