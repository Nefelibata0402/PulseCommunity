package main

import (
	"github.com/gin-gonic/gin"
	"newsCenter/cmd/config"
	"newsCenter/cmd/interfaces/rpc"
)

func initAll(r *gin.Engine) {
	initRouter(r)
	rpc.InitRpcUserClient()
}

func main() {
	r := gin.Default()
	initAll(r)
	err := r.Run(config.ApiConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
