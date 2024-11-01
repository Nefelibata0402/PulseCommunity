package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pulseCommunity/common/snowflake"
	"pulseCommunity/search/domain/event"
	"pulseCommunity/search/domain/service"
	"pulseCommunity/search/infrastructure/config"
	"pulseCommunity/search/infrastructure/persistence/mq"
)

func initAll() {
	config.InitConfig()
	//grpc服务注册
	RegisterGrpc()
	//grpc服务注册到etcd
	RegisterEtcdServer()
	err := snowflake.Init(1)
	if err != nil {
		zap.L().Error("snowflake.Init Fail", zap.Error(err))
	}
	saramaClient := mq.InitSaramaClient()
	userConsumer := event.NewUserConsumer(saramaClient, service.SyncServiceNew())
	articleConsumer := event.NewArticleConsumer(saramaClient, service.SyncServiceNew())
	consumers := mq.NewConsumers((*event.ArticleConsumer)(userConsumer), (*event.UserConsumer)(articleConsumer))
	for _, c := range consumers {
		err = c.Start()
		if err != nil {
			panic(err)
		}
	}
}

//func initPrometheus() {
//	go func() {
//		// 专门给 prometheus 用的端口
//		http.Handle("/metrics", promhttp.Handler())
//		http.ListenAndServe(":8081", nil)
//	}()
//}

func main() {
	r := gin.New()
	initAll()
	//initPrometheus()
	err := r.Run(config.SearchConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
