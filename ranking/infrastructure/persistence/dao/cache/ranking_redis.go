package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"pulseCommunity/ranking/domain/entity"
	"pulseCommunity/ranking/infrastructure/persistence/database/rediscache"
	"time"
)

type RankingRedis struct {
	rdb        *redis.Client
	key        string
	expiration time.Duration
}

func NewRankingRedis() *RankingRedis {
	return &RankingRedis{
		rdb:        rediscache.New(),
		key:        "ranking:top_n",
		expiration: time.Minute * 3,
	}
}
func (r *RankingRedis) Set(ctx context.Context, arts []entity.Article) error {
	for i := range arts {
		arts[i].Content = arts[i].Abstract()
	}
	val, err := json.Marshal(arts)
	if err != nil {
		return err
	}
	return r.rdb.Set(ctx, r.key, val, r.expiration).Err()
}

func (r *RankingRedis) Get(ctx context.Context) ([]entity.Article, error) {
	val, err := r.rdb.Get(ctx, r.key).Bytes()
	if err != nil {
		return nil, err
	}
	var res []entity.Article
	err = json.Unmarshal(val, &res)
	return res, err
}
