package main

import (
	"github.com/gin-gonic/gin"
	"newsCenter/common/snowflake"
	"newsCenter/user/application/router"
	"newsCenter/user/infrastructure/config"
)

func initAll(r *gin.Engine) {
	//grpc服务注册
	router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()
	snowflake.Init(1)
	config.InitConfig()
}

func main() {
	r := gin.Default()
	initAll(r)
	err := r.Run(config.UserConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
