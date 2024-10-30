package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"pulseCommunity/cmd/config"
	"pulseCommunity/cmd/interfaces/article"
	"pulseCommunity/cmd/interfaces/ranking"
	"pulseCommunity/cmd/interfaces/search"
	"pulseCommunity/cmd/interfaces/user/router"
	"pulseCommunity/cmd/middlewares/prometheus"
	"pulseCommunity/cmd/middlewares/ratelimit"
	"pulseCommunity/common/prometheus/new_prometheus"
	"time"
)

func initRouter(r *gin.Engine) {
	//限流
	redisClient := redis.NewClient(config.ApiConfig.ReadRedisConfig())
	builder := ratelimit.NewBuilder(redisClient, time.Second, 10)
	slideWindow := builder.Build()
	r.Use(slideWindow)
	pb := &prometheus.Builder{
		Namespace: "wang_cheng",
		Subsystem: "pulse_community",
		Name:      "gin_http",
		//InstanceId: "instanceId-1",
		Help: "统计GIN的HTTP接口数据",
	}
	pb1 := &prometheus.Builder{
		Namespace: "wang_cheng",
		Subsystem: "pulse_community",
		Name:      "current_active_requests",
		//InstanceId: "instanceId-2",
		Help: "统计当前请求活跃数量",
	}
	new_prometheus.InitPrometheus()

	r.Use(pb.BuildResponseTime())
	r.Use(pb1.BuildActiveRequest())
	router.InitUserRouter(r)
	article.InitArticleRouter(r)
	ranking.InitRankingRouter(r)
	search.InitSearchRouter(r)
}
