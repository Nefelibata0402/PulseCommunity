package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"newsCenter/article01/infrastructure/config"
)

var db *mongo.Database

func init() {
	username := config.ArticleConfig.MongoDBConfig.Username //账号
	password := config.ArticleConfig.MongoDBConfig.Password //密码
	host := config.ArticleConfig.MongoDBConfig.Host         //数据库地址，可以是Ip或者域名
	port := config.ArticleConfig.MongoDBConfig.Port         //数据库端口
	Dbname := config.ArticleConfig.MongoDBConfig.Db         //数据库名
	clientHostAndPort := fmt.Sprintf("mongodb://%s:%d", host, port)
	credential := options.Credential{
		Username: username,
		Password: password,
	}
	clientOptions := options.Client().ApplyURI(clientHostAndPort).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(Dbname)
}

func New() *mongo.Database {
	return db
}
