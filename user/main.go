package main

import (
	"github.com/gin-gonic/gin"
	"newsCenter/user/application/pkg/snowflake"
	"newsCenter/user/config"
	"newsCenter/user/router"
)

func initAll() {
	//initRouter(r)
	//grpc服务注册
	router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()
	snowflake.Init(1)
}

func main() {
	r := gin.Default()
	initAll()
	err := r.Run(config.UserConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
