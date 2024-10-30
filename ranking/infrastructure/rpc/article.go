package rpc

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"pulseCommunity/cmd/config"
	"pulseCommunity/common/discover"
	"pulseCommunity/idl/articleGrpc"
	"pulseCommunity/logs"
)

var ArticleServiceClient articleGrpc.ArticleServiceClient

func InitRpcArticleClient() {
	etcdRegister := discover.NewResolver(config.ApiConfig.EtcdConfig.Addr, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial("etcd:///article", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ArticleServiceClient = articleGrpc.NewArticleServiceClient(conn)
}

func GetArticleList(ctx context.Context, req *articleGrpc.GetArticleListRequest) (resp *articleGrpc.GetArticleListResponse, err error) {
	resp, err = ArticleServiceClient.GetArticleList(ctx, req)
	if err != nil {
		zap.L().Error("ArticleList GetArticleList Fail", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode != 200 {
		zap.L().Error("ArticleList GetArticleList Fail", zap.Error(err))
		return nil, errors.New("article ArticleList错误")
	}
	//得到的文章列表
	return resp, nil
}

func GetInteractiveByIds(ctx context.Context, req *articleGrpc.GetInteractiveByIdsRequest) (resp *articleGrpc.GetInteractiveByIdsResponse, err error) {
	resp, err = ArticleServiceClient.GetInteractiveByIds(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		zap.L().Error("ArticleList GetArticleList Fail", zap.Error(err))
		return nil, errors.New("article ArticleList错误")
	}
	//得到的交互列表
	return resp, nil
}
