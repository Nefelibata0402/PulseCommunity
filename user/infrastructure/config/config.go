package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"os"
	"pulseCommunity/logs"
)

var UserConfig = InitConfig()

type Config struct {
	Viper        *viper.Viper
	ServerConfig *ServerConfig
	GrpcConfig   *GrpcConfig
	EtcdConfig   *EtcdConfig
	MysqlConfig  *MysqlConfig
	JwtConfig    *JwtConfig
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
	Name    string
	Addr    string
	Version string
	Weight  int64
}

type EtcdConfig struct {
	Addr []string
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
}

type JwtConfig struct {
	AccessExp     int64
	RefreshExp    int64
	AccessSecret  string
	RefreshSecret string
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{Viper: v}
	workDir, _ := os.Getwd()
	conf.Viper.SetConfigName("config")
	conf.Viper.SetConfigType("yaml")
	conf.Viper.AddConfigPath(workDir + "/user/infrastructure/config")

	err := conf.Viper.ReadInConfig()
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
	conf.InitMysqlConfig()
	conf.InitJwtConfig()
	conf.ReadGrpcConfig()
	conf.InitJaegerConfig()
	conf.ReadRedisConfig()
	return conf
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.Viper.GetString("zap.debugFileName"),
		InfoFileName:  c.Viper.GetString("zap.infoFileName"),
		WarnFileName:  c.Viper.GetString("zap.warnFileName"),
		MaxSize:       c.Viper.GetInt("maxSize"),
		MaxAge:        c.Viper.GetInt("maxAge"),
		MaxBackups:    c.Viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.Viper.GetString("server.name")
	sc.Addr = c.Viper.GetString("server.addr")
	c.ServerConfig = sc
}

func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	var addr []string
	err := c.Viper.UnmarshalKey("etcd.addr", &addr)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addr = addr
	c.EtcdConfig = ec
}

func (c *Config) InitMysqlConfig() {
	mysql := &MysqlConfig{
		Username: c.Viper.GetString("mysql.username"),
		Password: c.Viper.GetString("mysql.password"),
		Host:     c.Viper.GetString("mysql.host"),
		Port:     c.Viper.GetInt("mysql.port"),
		Db:       c.Viper.GetString("mysql.db"),
	}
	c.MysqlConfig = mysql
}
func (c *Config) InitJwtConfig() {
	jwt := &JwtConfig{
		AccessSecret:  c.Viper.GetString("jwt.accessSecret"),
		AccessExp:     c.Viper.GetInt64("jwt.accessExp"),
		RefreshExp:    c.Viper.GetInt64("jwt.refreshExp"),
		RefreshSecret: c.Viper.GetString("jwt.refreshSecret"),
	}
	c.JwtConfig = jwt
}

func (c *Config) ReadGrpcConfig() {
	grpc := &GrpcConfig{}
	grpc.Name = c.Viper.GetString("grpc.name")
	grpc.Addr = c.Viper.GetString("grpc.addr")
	grpc.Version = c.Viper.GetString("grpc.version")
	grpc.Weight = c.Viper.GetInt64("grpc.weight")
	c.GrpcConfig = grpc
}

func (c *Config) InitJaegerConfig() {
	jc := &JaegerConfig{
		Endpoints: c.Viper.GetString("jaeger.endpoints"),
	}
	c.JaegerConfig = jc
}

func (c *Config) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     c.Viper.GetString("redis.host") + ":" + c.Viper.GetString("redis.port"),
		Password: c.Viper.GetString("redis.password"), // no password set
		DB:       c.Viper.GetInt("db"),                // user default DB
	}
}
