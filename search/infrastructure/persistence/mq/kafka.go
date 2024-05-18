package mq

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"newsCenter/search/domain/event"
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

func NewConsumers(articleConsumer *event.ArticleConsumer, userConsumer *event.UserConsumer) []event.Consumer {
	return []event.Consumer{
		articleConsumer,
		userConsumer,
	}
}
