package user

import (
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
