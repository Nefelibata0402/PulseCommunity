package gorms

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pulseCommunity/article/infrastructure/config"
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
	if err != nil {
		panic("数据库连接失败, error=" + err.Error())
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
