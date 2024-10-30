package service

import (
	"context"
	"golang.org/x/sync/errgroup"
	"pulseCommunity/search/domain/entity"
	"pulseCommunity/search/domain/repository"
	"pulseCommunity/search/infrastructure/persistence/dao"
	"strings"
)

type SearchService struct {
	userRepo    repository.UserRepository
	articleRepo repository.ArticleRepository
}

func SearchServiceNew() *SearchService {
	return &SearchService{
		userRepo:    dao.NewSyncElasticDao(),
		articleRepo: dao.NewSyncElasticDao(),
	}
}

type SearchServiceResponse interface {
	Search(ctx context.Context, expression string) (entity.SearchResult, error)
}

func (s *SearchService) Search(ctx context.Context, expression string) (entity.SearchResult, error) {
	// 你要搜索用户，你也要搜索 article
	// 要对 expression 进行解析，生成查询计划
	// 输入预处理
	// 清除掉空格，切割;',.
	keywords := strings.Split(expression, " ")
	var eg errgroup.Group
	var res entity.SearchResult
	eg.Go(func() error {
		users, err := s.userRepo.SearchUser(ctx, keywords)
		res.Users = users
		return err
	})
	eg.Go(func() error {
		arts, err := s.articleRepo.SearchArticle(ctx, keywords)
		res.Articles = arts
		return err
	})
	return res, eg.Wait()
}
