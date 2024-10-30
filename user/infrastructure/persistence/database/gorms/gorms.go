package gorms

import (
	"context"
	"fmt"
	prometheus2 "github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pulseCommunity/common/prometheus/gorm_prometheus"
	"pulseCommunity/user/infrastructure/config"
)

var _db *gorm.DB

func init() {
	//配置MySQL连接参数
	username := config.UserConfig.MysqlConfig.Username //账号
	password := config.UserConfig.MysqlConfig.Password //密码
	host := config.UserConfig.MysqlConfig.Host         //数据库地址，可以是Ip或者域名
	port := config.UserConfig.MysqlConfig.Port         //数据库端口
	Dbname := config.UserConfig.MysqlConfig.Db         //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//err = _db.Use(prometheus.New(prometheus.Config{
	//	DBName:          "pulse_community_user",
	//	RefreshInterval: 15,
	//	MetricsCollector: []prometheus.MetricsCollector{
	//		&prometheus.MySQL{
	//			VariableNames: []string{"thread_running"},
	//		},
	//	},
	//}))
	//if err != nil {
	//	panic(err)
	//}
	cb := gorm_prometheus.NewCallbacks(prometheus2.SummaryOpts{
		Namespace: "wang_cheng",
		Subsystem: "pulse_community",
		Name:      "gorm_db_user",
		Help:      "统计 GORM 的数据库查询",
		ConstLabels: map[string]string{
			"instance_id": "my_instance",
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

func (g *GormConn) Begin() {
	g.tx = GetDB().Begin()
}

func New() *GormConn {
	return &GormConn{db: GetDB()}
}

//	func NewTran() *GormConn {
//		return &GormConn{db: GetDB(), tx: GetDB()}
//	}
func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.db.Session(&gorm.Session{Context: ctx})
}

func (g *GormConn) Rollback() {
	g.tx.Rollback()
}
func (g *GormConn) Commit() {
	g.tx.Commit()
}

func (g *GormConn) Tx(ctx context.Context) *gorm.DB {
	return g.tx.WithContext(ctx)
}
