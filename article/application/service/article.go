package service

import (
	"context"
	"errors"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"pulseCommunity/article/domain/entity"
	articleEvent "pulseCommunity/article/domain/event/article"
	"pulseCommunity/article/domain/event/search"
	"pulseCommunity/article/domain/service"
	"pulseCommunity/article/infrastructure/persistence/mq"
	"pulseCommunity/article/infrastructure/rpc"
	"pulseCommunity/common/unierr"
	"pulseCommunity/idl/articleGrpc"
	"pulseCommunity/idl/userGrpc"
	"sync"
)

type ArticleService struct {
	articleGrpc.UnimplementedArticleServiceServer
	producerArticle articleEvent.Producer
	producerSearch  search.Producer
	biz             string
	mq              sarama.Client
	repo            service.ArticleServiceRepository
}

func New() *ArticleService {
	return &ArticleService{
		biz:             "article",
		mq:              mq.New(),
		producerArticle: articleEvent.NewSaramaSyncProducer(mq.InitSyncProducer(mq.New())),
		producerSearch:  search.NewSaramaSyncProducer(mq.InitSyncProducer(mq.New())),
		repo:            service.New(),
	}
}

func (article *ArticleService) Edit(c context.Context, req *articleGrpc.EditRequest) (resp *articleGrpc.EditResponse, err error) {
	//判断req.ArticleId是否大于0 大于0是更新 小于0是创建
	//articleService := service.New()
	art := entity.Article{
		Id:       req.ArticleId,
		Author:   entity.Author{Id: req.AuthorId},
		Content:  req.Data,
		Category: req.Category,
		Title:    req.Title,
	}
	if req.ArticleId > 0 {
		err = article.repo.Update(c, art)
		if err != nil {
			return nil, err
		}
	} else {
		err = article.repo.Create(c, art)
		if err != nil {
			zap.L().Error("Edit Create Fail", zap.Error(err))
			return nil, err
		}
	}
	resp = &articleGrpc.EditResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
	}
	return resp, nil
}
func (article *ArticleService) Publish(c context.Context, req *articleGrpc.PublishRequest) (resp *articleGrpc.PublishResponse, err error) {
	art := entity.Article{
		Id:       req.ArticleId,
		Author:   entity.Author{Id: req.AuthorId},
		Content:  req.Data,
		Category: req.Category,
		Title:    req.Title,
	}
	err = article.repo.Publish(c, art)
	if err != nil {
		zap.L().Error("Publish Publish Fail", zap.Error(err))
		return nil, err
	}
	go func() {
		err = article.producerSearch.ProduceReadEvent(search.ReadEvent{
			Id:      int64(req.ArticleId),
			Title:   req.Title,
			Status:  2,
			Content: req.Data,
		})
		if err != nil {
			zap.L().Error("Read ProduceReadEvent Fail 发送读事件失败", zap.Error(err))
			return
		}
	}()
	resp = &articleGrpc.PublishResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
	}
	return resp, nil
}

func (article *ArticleService) WithdrawArticle(c context.Context, req *articleGrpc.WithdrawRequest) (resp *articleGrpc.WithdrawResponse, err error) {
	err = article.repo.Withdraw(c, int64(req.AuthorId), int64(req.ArticleId))
	if err != nil {
		zap.L().Error("WithdrawArticle Withdraw Fail", zap.Error(err))
		return nil, err
	}
	resp = &articleGrpc.WithdrawResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
	}
	return resp, nil
}

func (article *ArticleService) GetDetail(c context.Context, req *articleGrpc.GetDetailRequest) (resp *articleGrpc.GetDetailResponse, err error) {
	//1.根据文章Id，获得文章详情
	art, err := article.repo.GetById(c, int64(req.ArticleId))
	if err != nil {
		zap.L().Error("GetDetail GetById Fail", zap.Error(err))
		return nil, err
	}
	//2.看看是不是作者
	if art.Author.Id != req.AuthorId {
		zap.L().Error("非法查询 次数多了 有人一直非法查询")
		return nil, errors.New("非法查询")
	}
	resp = &articleGrpc.GetDetailResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
		Article: &articleGrpc.Article{
			Id:       int64(art.Id),
			Author:   &userGrpc.User{Id: int64(art.Author.Id)},
			Content:  art.Content,
			Category: art.Category,
			Title:    art.Title,
			CreateAt: art.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt: art.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}
	return resp, nil
}

func (article *ArticleService) GetList(c context.Context, req *articleGrpc.GetListByAuthorRequest) (resp *articleGrpc.GetListByAuthorResponse, err error) {
	artList, err := article.repo.GetListByAuthor(c, int64(req.AuthorId), req.Offset, req.Limit)
	if err != nil {
		zap.L().Error("GetList GetListByAuthor Fail", zap.Error(err))
		return nil, err
	}
	List := make([]*articleGrpc.Article, 0)
	for i := 0; i < len(artList); i++ {
		List = append(List, &articleGrpc.Article{
			Id:       int64(artList[i].Id),
			Author:   &userGrpc.User{Id: int64(artList[i].Author.Id)},
			Content:  artList[i].Content,
			Category: artList[i].Category,
			Title:    artList[i].Title,
			CreateAt: artList[i].CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt: artList[i].UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	resp = &articleGrpc.GetListByAuthorResponse{
		StatusCode:  unierr.Success.ErrCode,
		StatusMsg:   unierr.Success.ErrMsg,
		ArticleList: List,
	}
	return resp, nil
}

func (article *ArticleService) Read(c context.Context, req *articleGrpc.ReadRequest) (resp *articleGrpc.ReadResponse, err error) {
	//1.获取文章信息-如果缓存中存在 直接返回
	var (
		errg        errgroup.Group
		flag        bool
		art         entity.Article
		interactive entity.Interactive
	)
	errg.Go(func() (errr error) {
		art, errr, flag = article.repo.GetArticleById(c, int64(req.ArticleId))
		//向消费者发送一条消息
		//别等 太费时了
		go func() {
			err = article.producerArticle.ProduceReadEvent(articleEvent.ReadEvent{
				ArticleId: int64(req.ArticleId),
				UserId:    int64(req.UserId),
			})
			if err != nil {
				zap.L().Error("Read ProduceReadEvent Fail 发送读事件失败", zap.Error(err))
				return
			}
		}()
		//2.未命中缓存
		if flag == false {
			//2.获取读者信息
			reqUserInfo := &userGrpc.UserInfoRequest{
				UserId: art.Author.Id,
			}
			res, errr := rpc.Info(c, reqUserInfo)
			if errr != nil {
				zap.L().Error("Read Info Fail", zap.Error(errr))
				return errr
			}
			//3.设置缓存
			art.Author.Name = res.User.Name
			errr = article.repo.SetPub(c, art)
			if errr != nil {
				zap.L().Error("Read SetPub Fail", zap.Error(errr))
				return errr
			}
		}
		return errr
	})
	errg.Go(func() (errr error) {
		//获取交互信息阅读计数 返回是否喜欢 是否收藏
		interactive, flag, errr = article.repo.GetInteractive(c, article.biz, int64(req.ArticleId), int64(req.UserId))
		if errr != nil {
			zap.L().Error("Read GetInteractive Fail", zap.Error(errr))
			return
		}
		if flag == false {
			var wg sync.WaitGroup
			wg.Add(2)
			var errrrg errgroup.Group
			errrrg.Go(func() (err error) {
				interactive.Liked, err = article.repo.Liked(c, article.biz, int64(req.ArticleId), int64(req.UserId))
				if err != nil {
					zap.L().Error("Read Liked Fail", zap.Error(err))
					return err
				}
				return nil
			})
			errrrg.Go(func() (err error) {
				interactive.Collected, err = article.repo.Collected(c, article.biz, int64(req.ArticleId), int64(req.UserId))
				if err != nil {
					zap.L().Error("Read Collecte Fail", zap.Error(err))
					return err
				}
				return nil
			})
			errr = errrrg.Wait()
			if errr != nil {
				return errr
			}
		}
		return
	})
	//改用Kafka
	//go func() {
	//	ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second)
	//	defer cannel()
	//	err = articleService.UpdateReadCnt(ctx, article.biz, int64(req.ArticleId))
	//	if err != nil {
	//		zap.L().Error("Read UpdateReadCnt Fail", zap.Error(err))
	//		return
	//	}
	//}()
	err = errg.Wait()
	if err != nil {
		zap.L().Error("Read Wait Fail", zap.Error(err))
		return nil, err
	}
	resp = &articleGrpc.ReadResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
		Article: &articleGrpc.Article{
			Id: int64(art.Id),
			Author: &userGrpc.User{
				Id:   int64(art.Author.Id),
				Name: art.Author.Name,
			},
			Content:    art.Content,
			Category:   art.Category,
			Title:      art.Title,
			ReadCnt:    interactive.ReadCnt,
			LikeCnt:    interactive.LikeCnt,
			CollectCnt: interactive.CollectCnt,
			Liked:      interactive.Liked,
			Collected:  interactive.Collected,
			CreateAt:   art.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:   art.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}
	return resp, nil
}

func (article *ArticleService) Like(c context.Context, req *articleGrpc.LikeRequest) (resp *articleGrpc.LikeResponse, err error) {
	//判断点赞还是取消点赞
	if req.Like == true {
		err = article.repo.Like(c, article.biz, req.Id, int64(req.UserId))
		if err != nil {
			return nil, err
		}
	} else {
		err = article.repo.CancelLike(c, article.biz, req.Id, int64(req.UserId))
		if err != nil {
			return nil, err
		}
	}
	resp = &articleGrpc.LikeResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
	}
	return resp, nil
}

// 获得文章列表 给热榜用
func (article *ArticleService) GetArticleList(ctx context.Context, req *articleGrpc.GetArticleListRequest) (resp *articleGrpc.GetArticleListResponse, err error) {
	artList, err := article.repo.GetList(ctx, req.StartTime, req.Offset, req.Limit)
	if err != nil {
		zap.L().Error("GetArticleList GetList Fail", zap.Error(err))
		return nil, err
	}
	List := make([]*articleGrpc.Article, 0)
	for i := 0; i < len(artList); i++ {
		List = append(List, &articleGrpc.Article{
			Id:       int64(artList[i].Id),
			Author:   &userGrpc.User{Id: int64(artList[i].Author.Id)},
			Content:  artList[i].Content,
			Category: artList[i].Category,
			Title:    artList[i].Title,
			CreateAt: artList[i].CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt: artList[i].UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	resp = &articleGrpc.GetArticleListResponse{
		StatusCode:  unierr.Success.ErrCode,
		StatusMsg:   unierr.Success.ErrMsg,
		ArticleList: List,
	}
	return resp, nil
}

// 获得交互列表 给热榜用
func (article *ArticleService) GetInteractiveByIds(c context.Context, req *articleGrpc.GetInteractiveByIdsRequest) (resp *articleGrpc.GetInteractiveByIdsResponse, err error) {
	interactiveList, err := article.repo.GetInteractiveByIds(c, req.Biz, req.IdsList)
	if err != nil {
		zap.L().Error("GetInteractiveByIds GetInteractiveByIds Fail", zap.Error(err))
		return nil, err
	}
	List := make([]*articleGrpc.Interactive, 0)
	for i := 0; i < len(interactiveList); i++ {
		List = append(List, &articleGrpc.Interactive{
			BizId:      interactiveList[i].BizId,
			ReadCnt:    interactiveList[i].ReadCnt,
			LikeCnt:    interactiveList[i].LikeCnt,
			CollectCnt: interactiveList[i].CollectCnt,
		})
	}
	resp = &articleGrpc.GetInteractiveByIdsResponse{
		StatusCode:      unierr.Success.ErrCode,
		StatusMsg:       unierr.Success.ErrMsg,
		InteractiveList: List,
	}
	return resp, nil
}
