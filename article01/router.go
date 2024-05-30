package main

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"newsCenter/article01/application/service"
	"newsCenter/article01/infrastructure/config"
	"newsCenter/common/discover"
	"newsCenter/idl/articleGrpc"
	"newsCenter/logs"
)

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	//0.0.0.0:8881
	c := gRPCConfig{
		Addr: config.ArticleConfig.GrpcConfig.Addr,
		RegisterFunc: func(g *grpc.Server) {
			articleGrpc.RegisterArticleServiceServer(g, service.New())
		}}
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
	//创建 etcd 服务发现解析器
	etcdRegister := discover.NewResolver(config.ArticleConfig.EtcdConfig.Addr, logs.LG)
	//etcd 解析器注册到 gRPC 框架中。这样，gRPC 将使用该解析器来解析服务地址。
	resolver.Register(etcdRegister)
	//服务地址:8881
	info := discover.Server{
		Name:    config.ArticleConfig.GrpcConfig.Name,
		Addr:    config.ArticleConfig.GrpcConfig.Addr,
		Version: config.ArticleConfig.GrpcConfig.Version,
		Weight:  config.ArticleConfig.GrpcConfig.Weight,
	}
	zap.L().Info("register grpc addr: ", zap.String("addr", info.Addr))
	//这行代码创建了一个新的注册器（register），用于将服务信息注册到 etcd 中。它也使用了 etcd 服务器的地址和一个日志记录器。
	r := discover.NewRegister(config.ArticleConfig.EtcdConfig.Addr, logs.LG)
	//调用了注册器的 Register 方法，将服务信息注册到 etcd 中
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
