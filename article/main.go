package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	articleEvent "newsCenter/article/domain/event/article"
	"newsCenter/article/domain/service"
	"newsCenter/article/infrastructure/config"
	"newsCenter/article/infrastructure/persistence/mq"
	"newsCenter/article/infrastructure/rpc"
	"newsCenter/common/snowflake"
)

func initAll(r *gin.Engine) {
	config.InitConfig()
	//grpc服务注册
	RegisterGrpc()
	//grpc服务注册到etcd
	RegisterEtcdServer()
	//注册用户rpc服务
	rpc.InitRpcUserClient()
	err := snowflake.Init(1)
	if err != nil {
		zap.L().Error("snowflake.Init Fail", zap.Error(err))
	}
	//consumer := articleEvent.NewInteractiveReadEventConsumer(service.New(), mq.New())
	//err = consumer.Start()
	//if err != nil {
	//	zap.L().Error("consumer.Start() Fail", zap.Error(err))
	//	return
	//}
	consumer := articleEvent.NewInteractiveReadEventConsumer(service.New(), mq.New())
	consumers := mq.InitConsumers(consumer)
	for _, c := range consumers {
		err = c.Start()
		if err != nil {
			zap.L().Error("consumer.Start() Fail", zap.Error(err))
			panic(err)
		}
	}
}

func main() {
	r := gin.New()
	//r.Use(logs.GinLogger())
	initAll(r)
	err := r.Run(config.ArticleConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
