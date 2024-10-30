package ranking

import (
	"github.com/gin-gonic/gin"
	"pulseCommunity/cmd/middlewares/tokenVerify"
)

func InitRankingRouter(r *gin.Engine) {
	router := r.Group("newsCenter/ranking")
	router.Use(tokenVerify.TokenVerify())
	router.GET("/TopN", TopN)
	router.GET("/getTopN", GetTopN) //查询热榜
}
