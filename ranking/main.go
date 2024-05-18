package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"newsCenter/common/snowflake"
	"newsCenter/ranking/application/job"
	"newsCenter/ranking/application/service"
	"newsCenter/ranking/infrastructure/config"
	"newsCenter/ranking/infrastructure/persistence/database/rediscache"
	"newsCenter/ranking/infrastructure/rpc"
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
	rpc.InitRpcArticleClient()
}

func main() {
	r := gin.New()
	initAll()
	cmdable := rediscache.InitRedis()
	rlockClient := rediscache.InitRlockClient(cmdable)
	//定时调度
	rankServer := job.InitRankingJob(service.New(), rlockClient)
	crron := job.InitJobs(rankServer)
	crron.Start()
	defer func() {
		// 等待定时任务退出
		<-crron.Stop().Done()
	}()
	err := r.Run(config.RankingConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
