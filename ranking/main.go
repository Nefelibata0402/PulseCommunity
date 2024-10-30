package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pulseCommunity/common/prometheus/new_prometheus"
	"pulseCommunity/common/snowflake"
	"pulseCommunity/ranking/application/job"
	"pulseCommunity/ranking/application/service"
	"pulseCommunity/ranking/infrastructure/config"
	"pulseCommunity/ranking/infrastructure/persistence/database/rediscache"
	"pulseCommunity/ranking/infrastructure/rpc"
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
	new_prometheus.InitPrometheusRanking()
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
