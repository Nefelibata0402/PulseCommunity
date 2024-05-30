package search

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

const topicSyncArticle = "sync_article_event"

type Producer interface {
	ProduceReadEvent(evt ReadEvent) error
}

type ReadEvent struct {
	Id      int64
	Title   string
	Status  int32
	Content string
}

type SaramaSyncProducer struct {
	producer sarama.SyncProducer
}

//	func NewSaramaSyncProducer(producer sarama.SyncProducer) Producer {
//		return &SaramaSyncProducer{producer: producer}
//	}
func NewSaramaSyncProducer(producer sarama.SyncProducer) Producer {
	return &SaramaSyncProducer{producer: producer}
}

func (s *SaramaSyncProducer) ProduceReadEvent(evt ReadEvent) error {
	val, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	_, _, err = s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topicSyncArticle,
		Value: sarama.StringEncoder(val),
	})
	zap.L().Info("生产者发送一条消息sync_article_event", zap.Any("读发送的值", evt))
	return err
}
