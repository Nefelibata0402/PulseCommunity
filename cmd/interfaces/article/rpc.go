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
	//用于将自定义的解析器注册到 gRPC 框架中
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial("etcd:///article",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{
  "loadBalancingConfig": [{"round_robin": {}}],
  "methodConfig":  [
    {
      "name": [{"service":  "ArticleService"}],
      "retryPolicy": {
        "maxAttempts": 4,
        "initialBackoff": "0.01s",
        "maxBackoff": "0.1s",
        "backoffMultiplier": 2.0,
        "retryableStatusCodes": ["UNAVAILABLE"]
      }
    }
  ]
}`),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ArticleServiceClient = articleGrpc.NewArticleServiceClient(conn)
}
