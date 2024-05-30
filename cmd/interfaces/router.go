package main

import (
	"github.com/gin-gonic/gin"
	"newsCenter/cmd/interfaces/article"
	"newsCenter/cmd/interfaces/ranking"
	"newsCenter/cmd/interfaces/search"
	"newsCenter/cmd/interfaces/user/router"
)

func initRouter(r *gin.Engine) {
	router.InitUserRouter(r)
	//限流
	//redisClient := redis.NewClient(config.ApiConfig.ReadRedisConfig())
	//builder := ratelimit.NewBuilder(redisClient, time.Second, 1)
	//slideWindow := builder.Build()
	//r.Use(slideWindow)
	article.InitArticleRouter(r)
	ranking.InitRankingRouter(r)
	search.InitSearchRouter(r)
}
