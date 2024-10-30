package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"pulseCommunity/cmd/config"
	"pulseCommunity/cmd/interfaces/article"
	"pulseCommunity/cmd/interfaces/ranking"
	"pulseCommunity/cmd/interfaces/search"
	"pulseCommunity/cmd/interfaces/user"
	"pulseCommunity/common/otel/init_otel"
	"time"
)

func initAll(r *gin.Engine) {
	initRouter(r)
	user.InitRpcUserClient()
	article.InitRpcArticleClient()
	ranking.InitRpcRankingClient()
	search.InitRpcSearchClient()
	config.InitConfig()
}

func main() {
	r := gin.Default()
	tpCancel := init_otel.InitOTEL()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		tpCancel(ctx)
	}()
	r.Use(otelgin.Middleware("middle_ware"))
	initAll(r)
	err := r.Run(config.ApiConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
