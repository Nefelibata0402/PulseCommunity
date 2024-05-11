package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"newsCenter/article/domain/entity"
	"newsCenter/article/domain/repository"
	"newsCenter/article/infrastructure/persistence/dao/cache"
	"newsCenter/article/infrastructure/persistence/dao/dao"
	"newsCenter/article/infrastructure/persistence/dao/mongodb"
	"time"
)

type ArticleService struct {
	articleRepo          repository.ArticleRepository      // mongodb
	cacheArticleRepo     repository.CacheArticleRepository // redis
	interactiveRepo      repository.InteractiveRepository  //mysql
	cacheInteractiveRepo repository.CacheInteractiveRepository
}

func New() *ArticleService {
	return &ArticleService{
		articleRepo:          mongodb.NewArticleMongoDB(),
		cacheArticleRepo:     cache.NewArticleRedis(),
		interactiveRepo:      dao.NewInteractiveDao(),
		cacheInteractiveRepo: cache.NewInteractiveRedis(),
	}
}

type ArticleServiceRepository interface {
	Update(c context.Context, art entity.Article) (err error)
	Create(c context.Context, art entity.Article) (err error)
	Publish(c context.Context, art entity.Article) (err error)
	Withdraw(ctx context.Context, uid int64, id int64) (err error)
	GetById(ctx context.Context, id int64) (art entity.Article, err error)
	GetListByAuthor(ctx context.Context, id int64, offset int64, limit int64) (artList []entity.Article, err error)
	preCache(ctx context.Context, arts []entity.Article)
	GetArticleById(ctx context.Context, id int64) (art entity.Article, err error, flag bool)
	SetPub(ctx context.Context, art entity.Article) (err error)
	GetInteractive(ctx context.Context, biz string, ArticleId int64, UserId int64) (interactive entity.Interactive, flag bool, err error)
	Liked(ctx context.Context, biz string, ArticleId int64, UserId int64) (flag bool, err error)
	Collected(ctx context.Context, biz string, ArticleId int64, UserId int64) (flag bool, err error)
	UpdateReadCnt(ctx context.Context, biz string, ArticleId int64) (err error)
	Like(c context.Context, biz string, Id int64, UserId int64) (err error)
	CancelLike(c context.Context, biz string, Id int64, UserId int64) (err error)
	BatchIncrReadCnt(ctx context.Context, biz []string, ArticleId []int64) error
}

func (article *ArticleService) Update(c context.Context, art entity.Article) (err error) {
	art.Status = entity.ArticleStatusNoPublish
	err = article.articleRepo.UpdateById(c, art)
	if err != nil {
		return err
	}
	err = article.cacheArticleRepo.DelFirstPage(c, int64(art.Author.Id))
	if err != nil {
		zap.L().Error("Create DelFirstPage Fail")
		return err
	}
	return err
}

func (article *ArticleService) Create(c context.Context, art entity.Article) (err error) {
	art.Status = entity.ArticleStatusNoPublish
	err = article.articleRepo.Insert(c, art)
	if err != nil {
		return err
	}
	err = article.cacheArticleRepo.DelFirstPage(c, int64(art.Author.Id))
	if err != nil {
		zap.L().Error("Create DelFirstPage Fail")
		return err
	}
	return err
}

func (article *ArticleService) Publish(c context.Context, art entity.Article) (err error) {
	art.Status = entity.ArticleStatusPublish
	err = article.articleRepo.Publish(c, art)
	if err != nil {
		zap.L().Error("Publish Publish Fail", zap.Error(err))
		return err
	}
	err = article.cacheArticleRepo.DelFirstPage(c, int64(art.Author.Id))
	if err != nil {
		zap.L().Error("Publish DelFirstPage Fail", zap.Error(err))
		return err
	}
	//设置缓存
	go func() {
		err = article.cacheArticleRepo.SetPub(c, art)
	}()
	return nil
}

func (article *ArticleService) Withdraw(ctx context.Context, uid int64, id int64) (err error) {
	err = article.articleRepo.Withdraw(ctx, uid, id, entity.ArticleStatusWithdraw)
	if err != nil {
		zap.L().Error("Withdraw Withdraw Fail", zap.Error(err))
		return err
	}
	err = article.cacheArticleRepo.DelFirstPage(ctx, uid)
	if err != nil {
		zap.L().Error("Withdraw DelFirstPage Fail", zap.Error(err))
		return err
	}
	return nil
}

func (article *ArticleService) GetById(ctx context.Context, id int64) (art entity.Article, err error) {
	art, err = article.cacheArticleRepo.Get(ctx, id)
	if err == nil {
		zap.L().Info("创作者查询文章细节走了缓存")
		return art, nil
	}
	art, err = article.articleRepo.GetById(ctx, id)
	if err != nil {
		zap.L().Error("GetById GetById Fail", zap.Error(err))
		return entity.Article{}, err
	}
	go func() {
		err = article.cacheArticleRepo.Set(ctx, art)
		if err != nil {
			zap.L().Error("GetById Set Fail", zap.Error(err))
			return
		}
	}()
	return art, nil
}

func (article *ArticleService) GetListByAuthor(ctx context.Context, id int64, offset int64, limit int64) (artList []entity.Article, err error) {
	//是否查询缓存
	if offset == 0 && limit == 10 {
		res, err := article.cacheArticleRepo.GetFirstPage(ctx, id)
		if err == nil {
			zap.L().Error("创作者查询文章列表走了缓存")
			return res, nil
		} else {
			zap.L().Error("创作者查询文章列表没查到缓存")
		}
	}
	//查询数据库
	artList, err = article.articleRepo.GetByAuthor(ctx, id, offset, limit)
	if err != nil {
		zap.L().Error("GetListByAuthor GetByAuthor Fail", zap.Error(err))
		return nil, err
	}
	//设置缓存
	go func() {
		if offset == 0 && limit == 10 {
			err = article.cacheArticleRepo.SetFirstPage(ctx, id, artList)
			if err != nil {
				zap.L().Error("GetListByAuthor SetFirstPage Fail", zap.Error(err))
			}
		}
		zap.L().Info("创作者查询文章列表 在没查到缓存的情况 查询了数据库 并设置了缓存")
	}()
	go func() {
		article.preCache(ctx, artList)
	}()
	return artList, nil
}

func (article *ArticleService) preCache(ctx context.Context, arts []entity.Article) {
	const size = 1024 * 1024
	if len(arts) > 0 && len(arts[0].Content) < size {
		err := article.cacheArticleRepo.Set(ctx, arts[0])
		if err != nil {
			zap.L().Error("preCache Set Fail", zap.Error(err))
		}
	}
}

func (article *ArticleService) GetArticleById(ctx context.Context, id int64) (art entity.Article, err error, flag bool) {
	res, err := article.cacheArticleRepo.GetPub(ctx, id)
	if err == nil {
		zap.L().Info("读者从缓存中得到了文章信息")
		return res, err, true
	}
	art, err = article.articleRepo.GetPubById(ctx, id)
	if err != nil {
		zap.L().Error("GetArticleById GetPubById Fail", zap.Error(err))
		return entity.Article{}, err, flag
	}
	return art, nil, flag
}
func (article *ArticleService) SetPub(ctx context.Context, art entity.Article) (err error) {
	err = article.cacheArticleRepo.SetPub(ctx, art)
	if err != nil {
		zap.L().Error("SetPub SetPub Fail", zap.Error(err))
		return err
	}
	zap.L().Info("读者查询文章列表 在没查到缓存的情况 查询了数据库 并设置了缓存")
	return nil
}

func (article *ArticleService) GetInteractive(ctx context.Context, biz string, ArticleId int64, UserId int64) (interactive entity.Interactive, flag bool, err error) {
	interactive, err = article.cacheInteractiveRepo.Get(ctx, biz, ArticleId)
	if err == nil {
		zap.L().Info("从交互接口缓存中得到交互信息")
		return interactive, true, nil
	}
	interactive, err = article.interactiveRepo.Get(ctx, biz, ArticleId)
	if err != nil {
		zap.L().Error("GetInteractive Get Fail", zap.Error(err))
		return entity.Interactive{}, false, err
	}
	if err == nil {
		err = article.cacheInteractiveRepo.Set(ctx, biz, ArticleId, interactive)
		if err != nil {
			zap.L().Error("GetInteractive Set Fail", zap.Error(err))
			return interactive, false, err
		}
		zap.L().Info("查询交互接口 在没查到缓存的情况 查询了数据库 并设置了缓存")
	}
	return interactive, false, nil
}

func (article *ArticleService) Liked(ctx context.Context, biz string, ArticleId int64, UserId int64) (flag bool, err error) {
	err = article.interactiveRepo.GetLikeInfo(ctx, biz, ArticleId, UserId)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	zap.L().Error("Like GetLikeInfo Fail", zap.Error(err))
	return false, err
}

func (article *ArticleService) Collected(ctx context.Context, biz string, ArticleId int64, UserId int64) (flag bool, err error) {
	err = article.interactiveRepo.GetCollectInfo(ctx, biz, ArticleId, UserId)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	zap.L().Error("Collect GetCollectInfo Fail", zap.Error(err))
	return false, err
}

func (article *ArticleService) UpdateReadCnt(ctx context.Context, biz string, ArticleId int64) (err error) {
	err = article.interactiveRepo.UpdateReadCnt(ctx, biz, ArticleId)
	if err != nil {
		zap.L().Error("UpdateReadCnt UpdateReadCnt Fail", zap.Error(err))
		return err
	}
	err = article.cacheInteractiveRepo.UpdateReadCntIfPresent(ctx, biz, ArticleId)
	if err != nil {
		zap.L().Error("UpdateReadCnt UpdateReadCntIfPresent Fail", zap.Error(err))
		return err
	}
	zap.L().Info("更新阅读计数 并更新交互接口总缓存")
	return nil
}

func (article *ArticleService) Like(c context.Context, biz string, Id int64, UserId int64) (err error) {
	err = article.interactiveRepo.InsertLikeInfo(c, biz, Id, UserId)
	if err != nil {
		zap.L().Error("Like InsertLikeInfo Fail", zap.Error(err))
		return err
	}
	err = article.cacheInteractiveRepo.UpdateLikeCntIfPresent(c, biz, Id)
	if err != nil {
		zap.L().Error("Like UpdateLikeCntIfPresent Fail", zap.Error(err))
		return err
	}
	zap.L().Info("增加点赞计数 并更新交互接口总缓存")
	return nil
}

func (article *ArticleService) CancelLike(c context.Context, biz string, Id int64, UserId int64) (err error) {
	err = article.interactiveRepo.DeleteLikeInfo(c, biz, Id, UserId)
	if err != nil {
		zap.L().Error("Like InsertLikeInfo Fail", zap.Error(err))
		return err
	}
	err = article.cacheInteractiveRepo.DeleteLikeCntIfPresent(c, biz, Id)
	if err != nil {
		zap.L().Error("Like UpdateLikeCntIfPresent Fail", zap.Error(err))
		return err
	}
	zap.L().Info("减少点赞计数 并更新交互接口总缓存")
	return nil
}

func (article *ArticleService) BatchIncrReadCnt(ctx context.Context, biz []string, ArticleId []int64) (err error) {
	err = article.interactiveRepo.BatchIncrReadCnt(ctx, biz, ArticleId)
	if err != nil {
		zap.L().Error("BatchIncrReadCnt BatchIncrReadCnt Fail", zap.Error(err))
		return err
	}
	go func() {
		c, cancel := context.WithTimeout(context.Background(), 180*time.Second)
		defer cancel()
		for i := 0; i < len(biz); i++ {
			err = article.cacheInteractiveRepo.UpdateReadCntIfPresent(c, biz[i], ArticleId[i])
			if err != nil {
				zap.L().Error("BatchIncrReadCnt UpdateReadCntIfPresent Fail", zap.Error(err))
			}
		}
	}()
	return nil
}
