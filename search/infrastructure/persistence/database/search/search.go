package search

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"newsCenter/search/infrastructure/config"
	"time"
)

var elasticClient *elastic.Client

func init() {
	const timeout = 3 * time.Second
	host := config.SearchConfig.ElasticSearchConfig.Host
	port := config.SearchConfig.ElasticSearchConfig.Port
	fmt.Println(host, port)
	clientHostAndPort := fmt.Sprintf("http://%s:%d", host, port)
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(clientHostAndPort),
		elastic.SetSniff(false),
		elastic.SetHealthcheckTimeoutStartup(timeout),
		elastic.SetTraceLog(log.Default()),
	}
	fmt.Println(host, port)
	client, err := elastic.NewClient(opts...)
	if err != nil {
		zap.L().Error("es Init Fail", zap.Error(err))
		panic(err)
	}
	err = InitES(client)
	if err != nil {
		panic(err)
	}
	elasticClient = client
	zap.L().Info("es启动成功")
}

func New() *elastic.Client {
	return elasticClient
}
