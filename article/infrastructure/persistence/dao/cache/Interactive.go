package cache

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/redis/go-redis/v9"
	"newsCenter/article/domain/entity"
	"newsCenter/article/infrastructure/persistence/database/rediscache"
	"strconv"
	"time"
)

type InteractiveRedis struct {
	rdb *redis.Client
}

func NewInteractiveRedis() *InteractiveRedis {
	return &InteractiveRedis{
		rdb: rediscache.New(),
	}
}

var (
	//go:embed incr_cnt.lua
	luaIncrCnt string
)

var ErrKeyNotExist = redis.Nil

const fieldReadCnt = "read_cnt"
const fieldLikeCnt = "like_cnt"
const fieldCollectCnt = "collect_cnt"

func (i *InteractiveRedis) Get(ctx context.Context, biz string, ArticleId int64) (inter entity.Interactive, err error) {
	key := i.key(biz, ArticleId)
	res, err := i.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return entity.Interactive{}, err
	}
	if len(res) == 0 {
		return entity.Interactive{}, ErrKeyNotExist
	}
	inter.CollectCnt, _ = strconv.ParseInt(res[fieldCollectCnt], 10, 64)
	inter.LikeCnt, _ = strconv.ParseInt(res[fieldLikeCnt], 10, 64)
	inter.ReadCnt, _ = strconv.ParseInt(res[fieldReadCnt], 10, 64)
	return inter, nil
}

func (i *InteractiveRedis) Set(ctx context.Context, biz string, ArticleId int64, inter entity.Interactive) (err error) {
	key := i.key(biz, ArticleId)
	err = i.rdb.HSet(ctx, key,
		fieldCollectCnt, inter.CollectCnt,
		fieldReadCnt, inter.ReadCnt,
		fieldLikeCnt, inter.LikeCnt,
	).Err()
	if err != nil {
		return err
	}
	err = i.rdb.Expire(ctx, key, time.Minute*15).Err()
	if err != nil {
		return err
	}
	return nil
}

func (i *InteractiveRedis) key(biz string, bizId int64) string {
	return fmt.Sprintf("interactive:%s:%d", biz, bizId)
}

func (i *InteractiveRedis) UpdateReadCntIfPresent(ctx context.Context, biz string, ArticleId int64) (err error) {
	key := i.key(biz, ArticleId)
	err = i.rdb.Eval(ctx, luaIncrCnt, []string{key}, fieldReadCnt, 1).Err()
	return err
}

func (i *InteractiveRedis) UpdateLikeCntIfPresent(ctx context.Context, biz string, id int64) (err error) {
	key := i.key(biz, id)
	err = i.rdb.Eval(ctx, luaIncrCnt, []string{key}, fieldLikeCnt, 1).Err()
	return err
}

func (i *InteractiveRedis) DeleteLikeCntIfPresent(ctx context.Context, biz string, id int64) (err error) {
	key := i.key(biz, id)
	//return fmt.Sprintf("interactive:%s:%d", biz, bizId) article:
	err = i.rdb.Eval(ctx, luaIncrCnt, []string{key}, fieldLikeCnt, -1).Err()
	return err
}
