package repository

import (
	"context"
	"pulseCommunity/article/domain/entity"
)

// ArticleRepository 文章接口MongoDB
type ArticleRepository interface {
	Insert(c context.Context, art entity.Article) error
	UpdateById(c context.Context, art entity.Article) error
	Publish(c context.Context, art entity.Article) error
	Withdraw(c context.Context, uid int64, id int64, status uint8) error
	GetById(ctx context.Context, id int64) (art entity.Article, err error)
	GetByAuthor(ctx context.Context, id int64, offset int64, limit int64) ([]entity.Article, error)
	GetList(ctx context.Context, startTime int64, offset int64, limit int64) ([]entity.Article, error)
	GetPubById(ctx context.Context, id int64) (art entity.Article, err error)
}

// CacheArticleRepository 文章接口Redis
type CacheArticleRepository interface {
	GetFirstPage(ctx context.Context, id int64) ([]entity.Article, error)
	SetFirstPage(ctx context.Context, id int64, artList []entity.Article) error
	DelFirstPage(ctx context.Context, id int64) error
	GetPub(ctx context.Context, id int64) (entity.Article, error)
	SetPub(ctx context.Context, article entity.Article) error
	Get(ctx context.Context, id int64) (entity.Article, error)
	Set(ctx context.Context, article entity.Article) error
}

// InteractiveRepository 交互接口MySQL
type InteractiveRepository interface {
	Get(ctx context.Context, biz string, articleId int64) (entity.Interactive, error)
	GetLikeInfo(ctx context.Context, biz string, ArticleId int64, UserId int64) error
	GetCollectInfo(ctx context.Context, biz string, ArticleId int64, UserId int64) error
	UpdateReadCnt(ctx context.Context, biz string, ArticleId int64) error
	InsertLikeInfo(ctx context.Context, biz string, ArticleId int64, UserId int64) error
	DeleteLikeInfo(ctx context.Context, biz string, ArticleId int64, UserId int64) error
	BatchIncrReadCnt(ctx context.Context, biz []string, ArticleId []int64) (err error)
	GetInteractiveByIds(ctx context.Context, biz string, ids []int64) ([]entity.Interactive, error)
}

// CacheInteractiveRepository 交互接口 Redis
type CacheInteractiveRepository interface {
	Get(ctx context.Context, biz string, ArticleId int64) (entity.Interactive, error)
	Set(ctx context.Context, biz string, bizId int64, inter entity.Interactive) error
	UpdateReadCntIfPresent(ctx context.Context, biz string, ArticleId int64) error
	UpdateLikeCntIfPresent(ctx context.Context, biz string, id int64) error
	DeleteLikeCntIfPresent(ctx context.Context, biz string, id int64) error
}
