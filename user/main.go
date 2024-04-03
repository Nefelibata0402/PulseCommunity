package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"newsCenter/user/application/pkg/snowflake"
	"newsCenter/user/application/trace"
	"newsCenter/user/config"
	"newsCenter/user/router"
)

func initAll(r *gin.Engine) {
	//initRouter(r)
	//grpc服务注册
	router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()
	snowflake.Init(1)
	tp, tpErr := trace.JaegerTraceProvider(config.UserConfig.JaegerConfig.Endpoints)
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}

func main() {
	r := gin.Default()
	initAll(r)
	err := r.Run(config.UserConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
