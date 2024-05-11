package rediscache

import (
	"github.com/redis/go-redis/v9"
	"newsCenter/article/infrastructure/config"
)

var Rc *redis.Client

func init() {
	rdb := redis.NewClient(config.ArticleConfig.ReadRedisConfig())
	Rc = rdb
}

func New() *redis.Client {
	return Rc
}
