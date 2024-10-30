package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"math"
	"pulseCommunity/common/unierr"
	"pulseCommunity/idl/articleGrpc"
	"pulseCommunity/idl/rankingGrpc"
	"pulseCommunity/idl/userGrpc"
	"pulseCommunity/ranking/domain/entity"
	"pulseCommunity/ranking/domain/service"
	"pulseCommunity/ranking/infrastructure/persistence/convertor"
	"pulseCommunity/ranking/infrastructure/pkg/pri_que/queue"
	"pulseCommunity/ranking/infrastructure/rpc"
	"time"
)

type RankingService struct {
	rankingGrpc.UnimplementedRankingServiceServer
	repo      service.RankingServiceRepository
	batchSize int64
	scoreFunc func(likeCnt int64, updatedTime time.Time) float64
	n         int
}

func New() *RankingService {
	return &RankingService{
		repo:      service.New(),
		batchSize: 100,
		n:         100,
		scoreFunc: func(likeCnt int64, updatedTime time.Time) float64 {
			// 时间
			duration := time.Since(updatedTime).Seconds()
			return float64(likeCnt-1) / math.Pow(duration+2, 1.5)
		},
	}
}

func (r *RankingService) GetTopN(ctx context.Context, req *rankingGrpc.GetTopNRequest) (resp *rankingGrpc.GetTopNResponse, err error) {
	res, err := r.repo.GetTopN(ctx)
	if err != nil {
		zap.L().Error("GetTopN GetTopN Fail", zap.Error(err))
		return nil, err
	}
	List := make([]*articleGrpc.Article, 0)
	for i := 0; i < len(res); i++ {
		List = append(List, &articleGrpc.Article{
			Id:       int64(res[i].Id),
			Author:   &userGrpc.User{Id: int64(res[i].Author.Id)},
			Content:  res[i].Content,
			Category: res[i].Category,
			Title:    res[i].Title,
			CreateAt: res[i].CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt: res[i].UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	resp = &rankingGrpc.GetTopNResponse{
		StatusCode:  unierr.Success.ErrCode,
		StatusMsg:   unierr.Success.ErrMsg,
		ArticleList: List,
	}
	return resp, nil
}

func (r *RankingService) TopN(ctx context.Context, req *rankingGrpc.TopNRequest) (resp *rankingGrpc.TopNResponse, err error) {
	zap.L().Info("TopN开始计算")
	arts, err := r.topN(ctx)
	if err != nil {
		zap.L().Error("TopN topN Fail", zap.Error(err))
		return nil, err
	}
	// 最终是要放到缓存里面的
	// 存到缓存里面
	err = r.repo.ReplaceTopN(ctx, arts)
	if err != nil {
		zap.L().Error("TopN ReplaceTopN Fail", zap.Error(err))
		return nil, err
	}
	resp = &rankingGrpc.TopNResponse{
		StatusCode: unierr.Success.ErrCode,
		StatusMsg:  unierr.Success.ErrMsg,
	}
	zap.L().Info("TopN计算结束")
	return resp, nil
}

//func (r *RankingService) TopN(ctx context.Context) error {
//	arts, err := r.topN(ctx)
//	if err != nil {
//		return err
//	}
//	// 最终是要放到缓存里面的
//	// 存到缓存里面
//	return r.repo.ReplaceTopN(ctx, arts)
//}

func (r *RankingService) topN(ctx context.Context) ([]entity.Article, error) {
	offset := 0
	//查询mongodb中时间
	start := time.Now().UnixMilli()
	//返回的时间
	start1 := time.Now()
	ddl := start1.Add(-7 * 24 * time.Hour)
	type Score struct {
		score float64
		art   entity.Article
	}
	topN := queue.NewPriorityQueue[Score](r.n, func(src Score, dst Score) int {
		if src.score > dst.score {
			return 1
		} else if src.score == dst.score {
			return 0
		} else {
			return -1
		}
	})
	//取文章数据
	for {
		//取文章数据 得到文章列表
		reqArticleList := &articleGrpc.GetArticleListRequest{
			StartTime: start,
			Offset:    int64(offset),
			Limit:     r.batchSize,
		}
		respArticleList, err := rpc.GetArticleList(ctx, reqArticleList)
		articleList := make([]entity.Article, 0)
		for _, val := range respArticleList.ArticleList {
			articleList = append(articleList, convertor.ArticleProtoToEntity(val))
		}
		if err != nil {
			zap.L().Error("TopN GetArticleList Fail", zap.Error(err))
			return nil, err
		}
		//提取出文章ids
		ids := make([]int64, 0)
		for _, val := range articleList {
			ids = append(ids, int64(val.Id))
		}
		//根据id取点赞数
		//从交互接口查询交互信息
		reqInteractive := &articleGrpc.GetInteractiveByIdsRequest{
			IdsList: ids,
			Biz:     "article",
		}
		respInteractiveList, err := rpc.GetInteractiveByIds(ctx, reqInteractive)
		if err != nil {
			zap.L().Error("TopN GetInteractiveByIds Fail", zap.Error(err))
			return nil, err
		}
		interactiveListMap := make(map[int64]entity.Interactive)
		for _, val := range respInteractiveList.InteractiveList {
			interactiveListMap[val.BizId] = convertor.InteractiveProtoToEntity(val)
		}
		for _, art := range respArticleList.ArticleList {
			//一条数据的点赞数等等
			inter := interactiveListMap[art.Id]
			//计算分数
			updatedTime, _ := time.Parse("2006-01-02 15:04:05", art.UpdateAt)
			score := r.scoreFunc(inter.LikeCnt, updatedTime)
			ele := Score{
				score: score,
				art:   convertor.ArticleProtoToEntity(art),
			}
			err = topN.Enqueue(ele)
			if err == queue.ErrOutOfCapacity {
				minEle, _ := topN.Dequeue()
				if minEle.score < score {
					_ = topN.Enqueue(ele)
				} else {
					_ = topN.Enqueue(ele)
				}
			}
		}
		offset = offset + len(articleList)
		if int64(len(articleList)) < r.batchSize || articleList[len(articleList)-1].UpdatedAt.Before(ddl) {
			break
		}
	}
	res := make([]entity.Article, r.n)
	//for i := 0; i < r.n; i++ {
	//	res[r.n-i-1] = heap.Pop(h).(Score).art
	//}
	for i := topN.Len() - 1; i >= 0; i-- {
		ele, _ := topN.Dequeue()
		res[i] = ele.art
	}
	fmt.Println(res)
	return res, nil
}

//func (r *RankingService) GetTopN(ctx context.Context) (res []entity.Article, err error) {
//	res, err = r.repo.GetTopN(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}

//func NewBatchRankingService(intrSvc InteractiveService, artSvc ArticleService) RankingService {
