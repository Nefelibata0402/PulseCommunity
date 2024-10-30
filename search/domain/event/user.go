package event

import (
	"context"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"pulseCommunity/search/domain/entity"
	"pulseCommunity/search/domain/service"
	saramax "pulseCommunity/search/infrastructure/pkg/samarax"
	"time"
)

const topicSyncUser = "sync_user_event"

type UserConsumer struct {
	syncSvc service.SyncServiceResponse
	client  sarama.Client
}

type UserEvent struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
}

func NewUserConsumer(client sarama.Client, svc service.SyncServiceResponse) *UserConsumer {
	return &UserConsumer{
		syncSvc: svc,
		client:  client,
	}
}

func (u *UserConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient("sync_user",
		u.client)
	if err != nil {
		return err
	}
	go func() {
		err := cg.Consume(context.Background(),
			[]string{topicSyncUser},
			saramax.NewHandler[UserEvent](u.Consume))
		if err != nil {
			zap.L().Error("user 退出了消费循环异常", zap.Error(err))
		}
	}()
	return err
}

func (u *UserConsumer) Consume(sg *sarama.ConsumerMessage, evt UserEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return u.syncSvc.InputUser(ctx, u.toEntity(evt))
}

func (u *UserConsumer) toEntity(evt UserEvent) entity.User {
	return entity.User{
		Id:       evt.Id,
		Nickname: evt.Nickname,
	}
}
