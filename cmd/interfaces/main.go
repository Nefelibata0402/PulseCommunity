package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"newsCenter/cmd/config"
	"newsCenter/cmd/interfaces/rpc"
	"newsCenter/cmd/trace"
)

func initAll(r *gin.Engine) {
	initRouter(r)
	rpc.InitRpcUserClient()
	tp, tpErr := trace.JaegerTraceProvider(config.ApiConfig.JaegerConfig.Endpoints)
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	r.Use(otelgin.Middleware("interfaces"))
}

func main() {
	r := gin.Default()
	initAll(r)
	err := r.Run(config.ApiConfig.ServerConfig.Addr)
	if err != nil {
		return
	}
}
