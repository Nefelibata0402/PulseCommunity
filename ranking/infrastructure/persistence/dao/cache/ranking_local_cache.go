package cache

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"newsCenter/ranking/domain/entity"
	"time"
)

type RankingLocalCache struct {
	topN       *Value[[]entity.Article]
	ddl        *Value[time.Time]
	expiration time.Duration
}

func NewRankingLocalCache() *RankingLocalCache {
	return &RankingLocalCache{
		topN:       NewValue[[]entity.Article](),
		ddl:        NewValue[time.Time](),
		expiration: 3 * time.Minute,
	}
}

func (r *RankingLocalCache) Set(ctx context.Context, arts []entity.Article) error {
	fmt.Println("调用Set初始化ddl")
	r.topN.Store(arts)
	r.ddl.Store(time.Now().Add(r.expiration))
	return nil
}

func (r *RankingLocalCache) Get(ctx context.Context) ([]entity.Article, error) {
	ddl := r.ddl.Load()
	arts := r.topN.Load()
	if len(arts) == 0 || ddl.Before(time.Now()) {
		zap.L().Error("Get 本地缓存失效")
		return nil, errors.New("本地缓存失效了")
	}
	return arts, nil
}

func (r *RankingLocalCache) ForceGet(ctx context.Context) ([]entity.Article, error) {
	arts := r.topN.Load()
	if len(arts) == 0 {
		zap.L().Error("ForceGet Load 本地缓存失效")
		return nil, errors.New("本地缓存失效了")
	}
	return arts, nil
}
