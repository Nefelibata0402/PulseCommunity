package service

import (
	"context"
	"go.uber.org/zap"
	"pulseCommunity/common/unierr"
	"pulseCommunity/idl/searchGrpc"
	"pulseCommunity/search/domain/service"
)

type SearchService struct {
	searchGrpc.UnimplementedSearchServiceServer
	repo service.SearchServiceResponse
	//prodecer articleEvent.Producer
	//mq       sarama.Client
	//repo     service.ArticleServiceRepository
}

func New() *SearchService {
	return &SearchService{
		repo: service.SearchServiceNew(),
		//mq:       mq.New(),
		//prodecer: articleEvent.NewSaramaSyncProducer(mq.InitSyncProducer(mq.New())),
		//repo:     service.New(),
	}
}

func (s SearchService) Search(c context.Context, req *searchGrpc.SearchRequest) (resp *searchGrpc.SearchResponse, err error) {
	res, err := s.repo.Search(c, req.Expression)
	if err != nil {
		zap.L().Error("Search Search Fail", zap.Error(err))
		return nil, err
	}
	userList := make([]*searchGrpc.User, 0)
	for i := 0; i < len(res.Users); i++ {
		userList = append(userList, &searchGrpc.User{
			Id:       resp.User[i].Id,
			Nickname: resp.User[i].Nickname,
		})
	}
	articleList := make([]*searchGrpc.Article, 0)
	for i := 0; i < len(res.Articles); i++ {
		articleList = append(articleList, &searchGrpc.Article{
			//Id:      res.Articles[i].Id,
			Title:   res.Articles[i].Title,
			Content: res.Articles[i].Content,
			//Status:  res.Articles[i].Status,
		})
	}
	resp = &searchGrpc.SearchResponse{
		User:       userList,
		Article:    articleList,
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
	}
	return resp, nil
}
