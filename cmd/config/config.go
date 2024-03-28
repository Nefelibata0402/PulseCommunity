package config

import "github.com/spf13/viper"

type Config struct {
	viper        *viper.Viper
	ServerConfig *ServerConfig
	GrpcConfig   *GrpcConfig
	EtcdConfig   *EtcdConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name string
	Addr string
}

type EtcdConfig struct {
	Addr []string
}
