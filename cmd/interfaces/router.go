package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"net/http"
	"newsCenter/cmd/config"
	"newsCenter/cmd/interfaces/article"
	"newsCenter/cmd/interfaces/ranking"
	"newsCenter/cmd/interfaces/search"
	"newsCenter/cmd/interfaces/user/router"
	prometheus "newsCenter/cmd/middlewares/prometheus"
	"newsCenter/cmd/middlewares/ratelimit"
	"time"
)

func initRouter(r *gin.Engine) {
	router.InitUserRouter(r)
	//限流
	redisClient := redis.NewClient(config.ApiConfig.ReadRedisConfig())
	builder := ratelimit.NewBuilder(redisClient, time.Second, 10)
	slideWindow := builder.Build()
	r.Use(slideWindow)
	pb := &prometheus.Builder{
		Namespace:  "wang_cheng",
		Subsystem:  "pulse_community",
		Name:       "gin_http",
		InstanceId: "instanceId-1",
		Help:       "统计GIN的HTTP接口数据",
	}
	pb1 := &prometheus.Builder{
		Namespace:  "wang_cheng",
		Subsystem:  "pulse_community",
		Name:       "current_active_requests",
		InstanceId: "instanceId-2",
		Help:       "统计当前请求活跃数量",
	}
	initPrometheus()
	r.Use(pb.BuildResponseTime())
	r.Use(pb1.BuildActiveRequest())
	article.InitArticleRouter(r)
	ranking.InitRankingRouter(r)
	search.InitSearchRouter(r)
}

func initPrometheus() {
	go func() {
		// 专门给 prometheus 用的端口
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()
}
