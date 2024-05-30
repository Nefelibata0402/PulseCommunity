package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"newsCenter/user/infrastructure/config"
	"time"
)

var Rc *RedisCache
var Cmd redis.Cmdable

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(config.UserConfig.ReadRedisConfig())
	Rc = &RedisCache{
		rdb: rdb,
	}
	Cmd = rdb
}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(ctx, key).Result()
	return result, err
}
