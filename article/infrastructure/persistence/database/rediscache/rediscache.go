package rediscache

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
	"pulseCommunity/article/infrastructure/config"
	"pulseCommunity/common/prometheus/redis_prometheus"
)

var Rc *redis.Client

func init() {
	rdb := redis.NewClient(config.ArticleConfig.ReadRedisConfig())
	Rc = rdb
	// 创建 Prometheus hook 并注册
	hook := redis_prometheus.NewPrometheusHook(prometheus.SummaryOpts{
		Namespace: "wang_cheng",
		Subsystem: "pulse_community",
		Name:      "article_redis",
		Help:      "统计 Redis 命中率",
		ConstLabels: map[string]string{
			"biz": "article",
		},
	})
	Rc.AddHook(hook)
	ctx := context.Background()
	if err := Rc.Set(ctx, "my_key", "my_value", 0).Err(); err != nil {
		fmt.Println("Set error:", err)
	}
	val, err := Rc.Get(ctx, "my_key").Result()
	if err != nil {
		fmt.Println("Get error:", err)
	} else {
		fmt.Println("Value:", val)
	}

	// 测试 Redis 操作
}

func New() *redis.Client {
	return Rc
}
