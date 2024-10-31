package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"pulseCommunity/logs"
)

type Config struct {
	Viper *viper.Viper
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{Viper: v}
	workDir, _ := os.Getwd()
	conf.Viper.SetConfigName("config")
	conf.Viper.SetConfigType("yaml")
	fmt.Println(workDir)
	conf.Viper.AddConfigPath(workDir + "/config")
	err := conf.Viper.ReadInConfig()
	if err != nil {
		log.Println("InitConfig test/config读取配置文件失败")
		return nil
	}
	conf.InitZapLog()
	zap.L().Info("日志器初始化成功")
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
		log.Fatalln("日志初始化失败", err)
	}
}
