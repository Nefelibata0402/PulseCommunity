package mq

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	eventArticle "newsCenter/article/domain/event/article"
)

var c sarama.Client

func init() {
	c = InitSaramaClient()
}

func InitSaramaClient() sarama.Client {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	client, err := sarama.NewClient([]string{"localhost:9094"}, saramaConfig)
	if err != nil {
		zap.L().Error("InitSaramaClient Fail")
		panic(err)
	}
	return client
}

func New() sarama.Client {
	return c
}

func InitSyncProducer(c sarama.Client) sarama.SyncProducer {
	p, err := sarama.NewSyncProducerFromClient(c)
	if err != nil {
		panic(err)
	}
	return p
}

func InitConsumers(c1 *eventArticle.InteractiveReadEventConsumer) []eventArticle.Consumer {
	return []eventArticle.Consumer{c1}
}
