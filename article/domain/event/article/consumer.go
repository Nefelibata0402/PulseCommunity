package article

import (
	"context"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"newsCenter/article/domain/service"
	"newsCenter/article/infrastructure/pkg/samarax"
	"time"
)

type InteractiveReadEventConsumer struct {
	//repo   service.ArticleService
	repo   service.ArticleServiceRepository
	client sarama.Client
}

func NewInteractiveReadEventConsumer(repo service.ArticleServiceRepository, client sarama.Client) *InteractiveReadEventConsumer {
	return &InteractiveReadEventConsumer{
		repo:   repo,
		client: client,
	}
}

func (i *InteractiveReadEventConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient("interactive", i.client)
	if err != nil {
		return err
	}
	go func() {
		er := cg.Consume(context.Background(),
			[]string{TopicReadEvent},
			samarax.NewBatchHandler[ReadEvent](i.BatchConsume))
		if er != nil {
			zap.L().Error("退出消费", zap.Error(err))
		}
	}()
	return err
}

func (i *InteractiveReadEventConsumer) BatchConsume(msgs []*sarama.ConsumerMessage,
	events []ReadEvent) error {
	bizs := make([]string, 0, len(events))
	bizIds := make([]int64, 0, len(events))
	for _, evt := range events {
		bizs = append(bizs, "article")
		bizIds = append(bizIds, evt.ArticleId)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	err := i.repo.BatchIncrReadCnt(ctx, bizs, bizIds)
	if err != nil {
		zap.L().Error("BatchConsume BatchIncrReadCnt Fail", zap.Error(err))
		return err
	}
	return nil
}

//func (i *InteractiveReadEventConsumer) Start() error {
//	cg, err := sarama.NewConsumerGroupFromClient("interactive", i.client)
//	if err != nil {
//		return err
//	}
//	go func() {
//		er := cg.Consume(context.Background(),
//			[]string{TopicReadEvent},
//			samarax.NewHandler[ReadEvent](i.Consume))
//		if er != nil {
//			zap.L().Error("退出消费", zap.Error(err))
//		}
//	}()
//	return err
//}

//func (i *InteractiveReadEventConsumer) Consume(msg *sarama.ConsumerMessage, event ReadEvent) error {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
//	defer cancel()
//	err := i.repo.UpdateReadCnt(ctx, "article", event.ArticleId)
//	if err != nil {
//		return err
//	}
//	zap.L().Info("消费者消费一条消息", zap.Any("消费的值", event))
//	return nil
//}
