package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"newsCenter/article/domain/entity"
	"newsCenter/article/infrastructure/persistence/database/rediscache"
	"time"
)

type ArticleRedis struct {
	rdb *redis.Client
}

func NewArticleRedis() *ArticleRedis {
	return &ArticleRedis{
		rdb: rediscache.New(),
	}
}

func (art *ArticleRedis) DelFirstPage(c context.Context, uid int64) (err error) {
	err = art.rdb.Del(c, art.firstKey(uid)).Err()
	zap.L().Info("删除首页缓存")
	return err
}

func (art *ArticleRedis) firstKey(uid int64) string {
	return fmt.Sprintf("article:first_page:%d", uid)
}

func (art *ArticleRedis) SetPub(ctx context.Context, article entity.Article) (err error) {
	val, err := json.Marshal(article)
	if err != nil {
		return err
	}
	err = art.rdb.Set(ctx, art.pubKey(int64(article.Id)), val, time.Minute*10).Err()
	if err != nil {
		return err
	}
	return err
}

func (art *ArticleRedis) pubKey(id int64) string {
	return fmt.Sprintf("article:pubish:detail:%d", id)
}

func (art *ArticleRedis) Get(ctx context.Context, id int64) (article entity.Article, err error) {
	val, err := art.rdb.Get(ctx, art.key(id)).Bytes()
	if err != nil {
		return entity.Article{}, err
	}
	var res entity.Article
	err = json.Unmarshal(val, &res)
	return res, err
}

func (art *ArticleRedis) key(id int64) string {
	return fmt.Sprintf("article:detail:%d", id)
}

func (art *ArticleRedis) Set(ctx context.Context, article entity.Article) (err error) {
	val, err := json.Marshal(article)
	if err != nil {
		return err
	}
	err = art.rdb.Set(ctx, art.key(int64(article.Id)), val, time.Minute*10).Err()
	if err != nil {
		return err
	}
	return nil
}

func (art *ArticleRedis) GetFirstPage(ctx context.Context, id int64) (articleList []entity.Article, err error) {
	key := art.firstKey(id)
	val, err := art.rdb.Get(ctx, key).Bytes()
	if err != nil {
		zap.L().Error("GetFirstPage Get Fail", zap.Error(err))
		return nil, err
	}
	err = json.Unmarshal(val, &articleList)
	if err != nil {
		zap.L().Error("GetFirstPage Unmarshal Fail", zap.Error(err))
		return nil, err
	}
	return articleList, nil
}

func (art *ArticleRedis) SetFirstPage(ctx context.Context, id int64, artList []entity.Article) error {
	for i := 0; i < len(artList); i++ {
		artList[i].Content = artList[i].Abstract()
	}
	key := art.firstKey(id)
	val, err := json.Marshal(artList)
	if err != nil {
		return err
	}
	err = art.rdb.Set(ctx, key, val, time.Minute*10).Err()
	if err != nil {
		return err
	}
	return nil
}

func (art *ArticleRedis) GetPub(ctx context.Context, id int64) (res entity.Article, err error) {
	val, err := art.rdb.Get(ctx, art.pubKey(id)).Bytes()
	if err != nil {
		return entity.Article{}, err
	}
	err = json.Unmarshal(val, &res)
	if err != nil {
		return entity.Article{}, err
	}
	return res, err
}
