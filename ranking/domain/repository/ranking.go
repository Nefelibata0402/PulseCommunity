package repository

import (
	"context"
	"newsCenter/ranking/domain/entity"
)

type RankingRedisCacheRepository interface {
	Set(ctx context.Context, arts []entity.Article) error
	Get(ctx context.Context) ([]entity.Article, error)
}

type RankingLocalCacheRepository interface {
	Set(ctx context.Context, arts []entity.Article) error
	Get(ctx context.Context) ([]entity.Article, error)
	ForceGet(ctx context.Context) ([]entity.Article, error)
}
