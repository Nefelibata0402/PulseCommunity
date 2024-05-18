package event

import (
	"context"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"newsCenter/search/domain/entity"
	"newsCenter/search/domain/service"
	saramax "newsCenter/search/infrastructure/pkg/samarax"
	"time"
)

const topicSyncArticle = "sync_article_event"

type ArticleConsumer struct {
	syncSvc service.SyncServiceResponse
	client  sarama.Client
}

func NewArticleConsumer(client sarama.Client, svc service.SyncServiceResponse) *ArticleConsumer {
	return &ArticleConsumer{
		syncSvc: svc,
		client:  client,
	}
}

type ArticleEvent struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Status  int32  `json:"status"`
	Content string `json:"content"`
}

func (a *ArticleConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient("sync_article", a.client)
	if err != nil {
		return err
	}
	go func() {
		err = cg.Consume(context.Background(),
			[]string{topicSyncArticle},
			saramax.NewHandler[ArticleEvent](a.Consume))
		if err != nil {
			zap.L().Error("article 退出了消费 循环异常", zap.Error(err))
		}
	}()
	return err
}

func (a *ArticleConsumer) Consume(sg *sarama.ConsumerMessage, evt ArticleEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	err := a.syncSvc.InputArticle(ctx, a.toEntity(evt))
	if err != nil {
		zap.L().Error("Consume InputArticle Fail", zap.Error(err))
		return err
	}
	return nil
}

func (a *ArticleConsumer) toEntity(article ArticleEvent) entity.Article {
	return entity.Article{
		Id:      article.Id,
		Title:   article.Title,
		Status:  article.Status,
		Content: article.Content,
	}
}
