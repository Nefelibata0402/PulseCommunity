package search

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"newsCenter/cmd/config"
	"newsCenter/common/discover"
	"newsCenter/idl/searchGrpc"
	"newsCenter/logs"
)

var SearchServiceClient searchGrpc.SearchServiceClient

func InitRpcSearchClient() {
	etcdRegister := discover.NewResolver(config.ApiConfig.EtcdConfig.Addr, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial("etcd:///search", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	SearchServiceClient = searchGrpc.NewSearchServiceClient(conn)
}
