package service

import (
	"context"
	"go.uber.org/zap"
	"pulseCommunity/search/domain/entity"
	"pulseCommunity/search/domain/repository"
	"pulseCommunity/search/infrastructure/persistence/dao"
)

type SyncService struct {
	userRepo    repository.UserRepository
	articleRepo repository.ArticleRepository
}

func SyncServiceNew() *SyncService {
	return &SyncService{
		userRepo:    dao.NewSyncElasticDao(),
		articleRepo: dao.NewSyncElasticDao(),
	}
}

type SyncServiceResponse interface {
	InputArticle(ctx context.Context, article entity.Article) error
	InputUser(ctx context.Context, user entity.User) error
}

func (s *SyncService) InputArticle(ctx context.Context, article entity.Article) error {
	err := s.articleRepo.InputArticle(ctx, article)
	if err != nil {
		zap.L().Error("InputArticle InputArticle Fail", zap.Error(err))
		return err
	}
	return nil
}

func (s *SyncService) InputUser(ctx context.Context, user entity.User) error {
	return s.userRepo.InputUser(ctx, user)
}
