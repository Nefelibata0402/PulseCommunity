package gorms

import (
	"context"
	"fmt"
	prometheus2 "github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pulseCommunity/article/infrastructure/config"
	"pulseCommunity/common/prometheus/gorm_prometheus"
)

var _db *gorm.DB

func init() {
	//配置MySQL连接参数
	username := config.ArticleConfig.MysqlConfig.Username //账号
	password := config.ArticleConfig.MysqlConfig.Password //密码
	host := config.ArticleConfig.MysqlConfig.Host         //数据库地址，可以是Ip或者域名
	port := config.ArticleConfig.MysqlConfig.Port         //数据库端口
	Dbname := config.ArticleConfig.MysqlConfig.Db         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	cb := gorm_prometheus.NewCallbacks(prometheus2.SummaryOpts{
		Namespace: "wang_cheng",
		Subsystem: "pulse_community",
		Name:      "article_gorm",
		Help:      "统计 article 的数据库查询",
		ConstLabels: map[string]string{
			"instance_id": "my_instance_gorm",
		},
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.75:  0.01,
			0.9:   0.01,
			0.99:  0.001,
			0.999: 0.0001,
		},
	})
	err = _db.Use(cb)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return _db
}

type GormConn struct {
	db *gorm.DB
	tx *gorm.DB
}

func New() *GormConn {
	return &GormConn{db: GetDB()}
}
func NewTran() *GormConn {
	return &GormConn{db: GetDB(), tx: GetDB()}
}
func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.db.Session(&gorm.Session{Context: ctx})
}
