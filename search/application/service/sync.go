package service

import (
	"context"
	"go.uber.org/zap"
	"newsCenter/idl/searchGrpc"
	"newsCenter/search/domain/entity"
	"newsCenter/search/domain/service"
)

type SyncService struct {
	searchGrpc.UnimplementedSyncServiceServer
	repo service.SyncServiceResponse
}

func SyncServiceNew() *SyncService {
	return &SyncService{
		repo: service.SyncServiceNew(),
	}
}

func (s *SyncService) InputUser(c context.Context, req *searchGrpc.InputUserRequest) (resp *searchGrpc.InputUserResponse, err error) {
	err = s.repo.InputUser(c, s.toEntityUser(req.User))
	if err != nil {
		zap.L().Error("InputUser InputUser Fail", zap.Error(err))
		return nil, err
	}
	return resp, err
}
func (s *SyncService) InputArticle(c context.Context, req *searchGrpc.InputArticleRequest) (resp *searchGrpc.InputArticleResponse, err error) {
	err = s.repo.InputArticle(c, s.toEntityArticle(req.Article))
	if err != nil {
		zap.L().Error("InputArticle InputArticle Fail", zap.Error(err))
		return nil, err
	}
	return resp, err
}

func (s *SyncService) toEntityUser(user *searchGrpc.User) entity.User {
	return entity.User{
		Id:       user.Id,
		Nickname: user.Nickname,
	}
}

func (s *SyncService) toEntityArticle(art *searchGrpc.Article) entity.Article {
	return entity.Article{
		Id:      art.Id,
		Title:   art.Title,
		Status:  art.Status,
		Content: art.Content,
	}
}
