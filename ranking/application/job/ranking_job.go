package job

import (
	"context"
	"go.uber.org/zap"
	"newsCenter/idl/rankingGrpc"
	"newsCenter/ranking/infrastructure/pkg/redis_lock"
	"sync"
	"time"
)

type RankingJob struct {
	svc       rankingGrpc.RankingServiceServer
	timeout   time.Duration
	client    *redis_lock.Client
	lock      *redis_lock.Lock
	key       string
	localLock *sync.Mutex
	load      int32
}

func NewRankingJob(svc rankingGrpc.RankingServiceServer, client *redis_lock.Client, timeout time.Duration) *RankingJob {
	return &RankingJob{
		svc:       svc,
		timeout:   timeout,
		key:       "job:ranking",
		client:    client,
		localLock: &sync.Mutex{},
	}
}

func (r *RankingJob) Name() string {
	return "ranking"
}

func (r *RankingJob) Run() error {
	//r.localLock.Lock()
	//lock := r.lock
	//if lock == nil {
	//	//抢分布式锁
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	//	defer cancel()
	//	lock, err := r.client.Lock(ctx, r.key, r.timeout, &redis_lock.FixIntervalRetry{
	//		Interval: time.Millisecond * 100,
	//		Max:      5,
	//		//重试的超时
	//	}, time.Second)
	//	if err != nil {
	//		r.localLock.Unlock()
	//		zap.L().Info("获取分布式锁失败", zap.Error(err))
	//		return nil
	//	}
	//	r.lock = lock
	//	r.localLock.Unlock()
	//	go func() {
	//		// 并不是非得一半就续约
	//		err = lock.AutoRefresh(r.timeout/2, r.timeout)
	//		if err != nil {
	//			// 续约失败了
	//			// 你也没办法中断当下正在调度的热榜计算（如果有）
	//			r.localLock.Lock()
	//			r.lock = nil
	//			r.localLock.Unlock()
	//		}
	//	}()
	//}
	//拿到了锁
	//zap.L().Info("拿到了分布式锁")
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	req := &rankingGrpc.TopNRequest{}
	_, err := r.svc.TopN(ctx, req)
	if err != nil {
		zap.L().Error("Run TopN", zap.Error(err))
		return err
	}
	return nil
}

func (r *RankingJob) Close() error {
	r.localLock.Lock()
	lock := r.lock
	r.localLock.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return lock.Unlock(ctx)
}
