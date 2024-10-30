package main

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	discover2 "pulseCommunity/common/discover"
	"pulseCommunity/idl/rankingGrpc"
	"pulseCommunity/logs"
	"pulseCommunity/ranking/application/service"
	"pulseCommunity/ranking/infrastructure/config"
)

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	//0.0.0.0:8881
	c := gRPCConfig{
		Addr: config.RankingConfig.GrpcConfig.Addr,
		RegisterFunc: func(g *grpc.Server) {
			rankingGrpc.RegisterRankingServiceServer(g, service.New())
		},
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		otelgrpc.UnaryServerInterceptor(),
	)))
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		log.Printf("grpc server started as: %s \n", c.Addr)
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

func RegisterEtcdServer() {
	etcdRegister := discover2.NewResolver(config.RankingConfig.EtcdConfig.Addr, logs.LG)
	resolver.Register(etcdRegister)
	//服务地址:8881
	info := discover2.Server{
		Name:    config.RankingConfig.GrpcConfig.Name,
		Addr:    config.RankingConfig.GrpcConfig.Addr,
		Version: config.RankingConfig.GrpcConfig.Version,
		Weight:  config.RankingConfig.GrpcConfig.Weight,
	}
	zap.L().Info("register grpc addr: ", zap.String("addr", info.Addr))
	r := discover2.NewRegister(config.RankingConfig.EtcdConfig.Addr, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
