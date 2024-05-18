package service

import (
	"context"
	"go.uber.org/zap"
	"newsCenter/ranking/domain/entity"
	"newsCenter/ranking/domain/repository"
	"newsCenter/ranking/infrastructure/persistence/dao/cache"
)

type RankingService struct {
	//repo  repository.RankingRepository
	redisCache repository.RankingRedisCacheRepository
	localCache repository.RankingLocalCacheRepository
}

func New() *RankingService {
	return &RankingService{
		redisCache: cache.NewRankingRedis(),
		localCache: cache.NewRankingLocalCache(),
	}
}

type RankingServiceRepository interface {
	ReplaceTopN(ctx context.Context, arts []entity.Article) error
	GetTopN(ctx context.Context) ([]entity.Article, error)
}

func (r *RankingService) ReplaceTopN(ctx context.Context, arts []entity.Article) (err error) {
	err = r.localCache.Set(ctx, arts)
	if err != nil {
		zap.L().Error("ReplaceTopN Set Fail 本地缓存设置失败", zap.Error(err))
	}
	zap.L().Info("成功设置本地缓存")
	err = r.redisCache.Set(ctx, arts)
	if err != nil {
		zap.L().Error("ReplaceTopN Set Fail", zap.Error(err))
		return err
	}
	zap.L().Info("成功设置redis缓存")
	return nil
}

func (r *RankingService) GetTopN(ctx context.Context) (arts []entity.Article, err error) {
	arts, err = r.localCache.Get(ctx)
	if err == nil {
		zap.L().Info("从本地缓存中获取数据")
		return arts, nil
	}
	//回写本地缓存
	arts, err = r.redisCache.Get(ctx)
	if err == nil {
		_ = r.localCache.Set(ctx, arts)
	} else {
		return r.localCache.ForceGet(ctx)
	}
	return arts, nil
}

//func (r *RankingService) GetTopN(ctx context.Context) (arts []entity.Article, err error) {
//	arts, err = r.redisCache.Get(ctx)
//	if err != nil {
//		zap.L().Error("GetTopN Get Fail", zap.Error(err))
//		return nil, err
//	}
//	return arts, nil
//}
