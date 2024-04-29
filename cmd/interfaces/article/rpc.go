package article

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"newsCenter/cmd/config"
	"newsCenter/common/discover"
	"newsCenter/idl/articleGrpc"
	"newsCenter/logs"
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
