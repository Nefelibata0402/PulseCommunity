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
	"pulseCommunity/idl/userGrpc"
	"pulseCommunity/logs"
)

var UserServiceClient userGrpc.UserServiceClient

func InitRpcUserClient() {
	etcdRegister := discover.NewResolver(config.ApiConfig.EtcdConfig.Addr, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserServiceClient = userGrpc.NewUserServiceClient(conn)
}

func Info(ctx context.Context, req *userGrpc.UserInfoRequest) (resp *userGrpc.UserInfoResponse, err error) {
	resp, err = UserServiceClient.GetUserinfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		zap.L().Error("Info GetUserinfo Fail 调用User rpc错误", zap.Error(err))
		return nil, errors.New("调用User rpc错误")
	}
	return resp, nil
}
