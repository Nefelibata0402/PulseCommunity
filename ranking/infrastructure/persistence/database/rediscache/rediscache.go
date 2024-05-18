package rediscache

import (
	"github.com/redis/go-redis/v9"
	"newsCenter/ranking/infrastructure/config"
	"newsCenter/ranking/infrastructure/pkg/redis_lock"
)

var Rc *redis.Client

func init() {
	rdb := redis.NewClient(config.RankingConfig.ReadRedisConfig())
	Rc = rdb
}

func New() *redis.Client {
	return Rc
}

func InitRedis() redis.Cmdable {
	return redis.NewClient(&redis.Options{Addr: config.RankingConfig.ReadRedisConfig().Addr})
}

func InitRlockClient(client redis.Cmdable) *redis_lock.Client {
	return redis_lock.NewClient(client)
}
