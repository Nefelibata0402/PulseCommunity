package article

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

const TopicReadEvent = "article_read"

type Producer interface {
	ProduceReadEvent(evt ReadEvent) error
}

type ReadEvent struct {
	ArticleId int64
	UserId    int64
}

type BatchReadEvent struct {
	ArticleId []int64
	UserId    []int64
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
		Topic: TopicReadEvent,
		Value: sarama.StringEncoder(val),
	})
	zap.L().Info("生产者发送一条消息", zap.Any("读发送的值", evt))
	return err
}
