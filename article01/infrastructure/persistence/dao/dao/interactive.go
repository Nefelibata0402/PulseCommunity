package dao

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"newsCenter/article01/domain/entity"
	"newsCenter/article01/infrastructure/persistence/convertor"
	"newsCenter/article01/infrastructure/persistence/database/collection"
	"newsCenter/article01/infrastructure/persistence/database/gorms"
	"newsCenter/article01/infrastructure/persistence/database/interactive"
	"newsCenter/article01/infrastructure/persistence/database/like"
	"time"
)

type InteractiveGorm struct {
	conn *gorms.GormConn
	tran *Transaction
}

func NewInteractiveDao() *InteractiveGorm {
	return &InteractiveGorm{
		conn: gorms.New(),
		tran: NewTransaction(),
	}
}

func (i *InteractiveGorm) Get(ctx context.Context, biz string, articleId int64) (inter entity.Interactive, err error) {
	var res interactive.Interactive
	err = i.conn.Session(ctx).Where("biz = ? AND biz_id = ?", biz, articleId).First(&res).Error
	if err != nil {
		return entity.Interactive{}, err
	}
	inter = convertor.ToInteractiveEntity(res)
	return inter, nil
}

func (i *InteractiveGorm) GetLikeInfo(ctx context.Context, biz string, ArticleId int64, UserId int64) (err error) {
	var res like.Like
	err = i.conn.Session(ctx).Where("biz = ? AND biz_id = ? AND id = ? AND status = ?", biz, ArticleId, UserId, 1).First(&res).Error
	return err
}

func (i *InteractiveGorm) GetCollectInfo(ctx context.Context, biz string, ArticleId int64, UserId int64) (err error) {
	var res collection.Collection
	err = i.conn.Session(ctx).Where("biz = ? AND biz_id = ? AND id = ?", biz, ArticleId, UserId).First(&res).Error
	return err
}

func (i *InteractiveGorm) UpdateReadCnt(ctx context.Context, biz string, ArticleId int64) (err error) {
	now := time.Now().UnixMilli()
	err = i.conn.Session(ctx).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{
			"read_cnt":   gorm.Expr("`read_cnt` + 1"),
			"updated_at": now,
		}),
	}).Create(&interactive.Interactive{
		Biz:       biz,
		BizId:     ArticleId,
		ReadCnt:   1,
		CreatedAt: now,
		UpdatedAt: now,
	}).Error
	return err
}

func (i *InteractiveGorm) InsertLikeInfo(ctx context.Context, biz string, ArticleId int64, UserId int64) (err error) {
	now := time.Now().UnixMilli()
	var res like.Like
	//防止重复点赞
	err = i.conn.Session(ctx).Where("biz = ? AND biz_id = ? AND id = ?", biz, ArticleId, UserId).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || res.Status == 0 {
		err = i.tran.Action(func(conn gorms.DbConn) error {
			i.conn = conn.(*gorms.GormConn)
			err = i.conn.Tx(ctx).Clauses(clause.OnConflict{
				DoUpdates: clause.Assignments(map[string]interface{}{
					"updated_at": now,
					"status":     1,
				}),
			}).Create(&like.Like{
				Id:        UserId,
				BizId:     ArticleId,
				Biz:       biz,
				Status:    1,
				CreatedAt: now,
				UpdatedAt: now,
			}).Error
			if err != nil {
				return err
			}
			err = i.conn.Tx(ctx).Clauses(clause.OnConflict{
				DoUpdates: clause.Assignments(map[string]interface{}{
					"like_cnt":   gorm.Expr("`like_cnt` + 1"),
					"updated_at": now,
				}),
			}).Create(&interactive.Interactive{
				BizId:     ArticleId,
				Biz:       biz,
				LikeCnt:   1,
				CreatedAt: now,
				UpdatedAt: now,
			}).Error
			if err != nil {
				return err
			}
			return nil
		})
		return err
	}
	if res.Status == 1 {
		zap.L().Info("重复点赞")
		return errors.New("重复点赞")
	}
	return nil
}

func (i *InteractiveGorm) DeleteLikeInfo(ctx context.Context, biz string, ArticleId int64, UserId int64) (err error) {
	now := time.Now().UnixMilli()
	//防止重复删除
	var res like.Like
	err = i.conn.Session(ctx).Where("biz = ? AND biz_id = ? AND id = ?", biz, ArticleId, UserId).First(&res).Error
	if res.Status == 0 {
		zap.L().Info("重复删除")
		return errors.New("重复删除")
	}
	err = i.tran.Action(func(conn gorms.DbConn) error {
		i.conn = conn.(*gorms.GormConn)
		err = i.conn.Tx(ctx).Model(&like.Like{}).
			Where("id = ? AND biz_id = ? AND biz = ?", UserId, ArticleId, biz).
			Updates(map[string]interface{}{
				"updated_at": now,
				"status":     0,
			}).Error
		if err != nil {
			return err
		}
		err = i.conn.Tx(ctx).Model(&interactive.Interactive{}).
			Where("biz = ? AND biz_id = ?", biz, ArticleId).
			Updates(map[string]interface{}{
				"like_cnt":   gorm.Expr("`like_cnt` - 1"),
				"updated_at": now,
			}).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (i *InteractiveGorm) BatchIncrReadCnt(ctx context.Context, biz []string, ArticleId []int64) (err error) {
	err = i.conn.Session(ctx).Transaction(func(tx *gorm.DB) error {
		for x := 0; x < len(biz); x++ {
			err = i.UpdateReadCnt(ctx, biz[x], ArticleId[x])
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (i *InteractiveGorm) GetInteractiveByIds(ctx context.Context, biz string, ids []int64) (res []entity.Interactive, err error) {
	var tmp []interactive.Interactive
	err = i.conn.Session(ctx).Where("biz = ? AND biz_id IN ?", biz, ids).First(&tmp).Error
	if err != nil {
		return nil, err
	}
	for _, val := range tmp {
		res = append(res, convertor.ToInteractiveEntity(val))
	}
	return res, nil
}
