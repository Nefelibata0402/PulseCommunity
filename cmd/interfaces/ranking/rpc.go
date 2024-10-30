package ranking

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"pulseCommunity/cmd/config"
	"pulseCommunity/common/discover"
	"pulseCommunity/idl/rankingGrpc"
	"pulseCommunity/logs"
)

var RankingServiceClient rankingGrpc.RankingServiceClient

func InitRpcRankingClient() {
	etcdRegister := discover.NewResolver(config.ApiConfig.EtcdConfig.Addr, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial("etcd:///ranking", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	RankingServiceClient = rankingGrpc.NewRankingServiceClient(conn)
}
