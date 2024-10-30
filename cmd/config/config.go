package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"pulseCommunity/logs"
)

var ApiConfig = InitConfig()

type Config struct {
	viper        *viper.Viper
	ServerConfig *ServerConfig
	GrpcConfig   *GrpcConfig
	EtcdConfig   *EtcdConfig
	JaegerConfig *JaegerConfig
}
type JaegerConfig struct {
	Endpoints string
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

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	//workDir, _ := os.Getwd()
	//conf.viper.SetConfigName("config")
	//conf.viper.SetConfigType("yaml")
	//conf.viper.AddConfigPath(workDir + "/cmd/config")
	conf.viper.AddConfigPath("/Users/wangcheng/Documents/golang/src/pulseCommunity/cmd/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.ReadEtcdConfig()
	conf.InitJaegerConfig()
	conf.ReadRedisConfig()
	return conf
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.ServerConfig = sc
}

func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	var addr []string
	err := c.viper.UnmarshalKey("etcd.addr", &addr)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addr = addr
	c.EtcdConfig = ec
}

func (c *Config) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("db"),
	}
}

func (c *Config) InitJaegerConfig() {
	jc := &JaegerConfig{
		Endpoints: c.viper.GetString("jaeger.endpoints"),
	}
	c.JaegerConfig = jc
}
