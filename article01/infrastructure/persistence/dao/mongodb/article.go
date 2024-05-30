package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"newsCenter/article01/domain/entity"
	"newsCenter/article01/infrastructure/persistence/convertor"
	artil "newsCenter/article01/infrastructure/persistence/database/article"
	"newsCenter/article01/infrastructure/persistence/database/mongodb"
	"newsCenter/common/snowflake"
	"time"
)

type ArticleMongoDB struct {
	col    *mongo.Collection
	pubCol *mongo.Collection
}

func NewArticleMongoDB() *ArticleMongoDB {
	return &ArticleMongoDB{
		col:    mongodb.New().Collection("articles"),
		pubCol: mongodb.New().Collection("published_articles"),
	}
}

func (article *ArticleMongoDB) Insert(c context.Context, art entity.Article) (err error) {
	artDao := convertor.ToDao(art)
	now := time.Now().UnixMilli()
	artDao.CreatedAt = now
	artDao.UpdatedAt = now
	artId, err := snowflake.GetID()
	if err != nil {
		return err
	}
	artDao.Id = artId
	_, err = article.col.InsertOne(c, &artDao)
	if err != nil {
		return err
	}
	return nil
}

func (article *ArticleMongoDB) UpdateById(c context.Context, art entity.Article) (err error) {
	artDao := convertor.ToDao(art)
	now := time.Now().UnixMilli()
	filter := bson.D{bson.E{Key: "id", Value: artDao.Id},
		bson.E{Key: "user_id", Value: artDao.UserId}}
	set := bson.D{bson.E{Key: "$set", Value: bson.M{
		"title":      artDao.Title,
		"content":    artDao.Content,
		"status":     artDao.Status,
		"updated_at": now,
	}}}
	res, err := article.col.UpdateOne(c, filter, set)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		zap.L().Error("创作者ID不正确")
		return errors.New("你不是作者（ID不正确）")
	}
	return nil
}

func (article *ArticleMongoDB) Publish(c context.Context, art entity.Article) (err error) {
	artDao := convertor.ToDao(art)
	//防止并发问题 手机电脑同时编辑 有一个先放弃了 另一个后发布 先插入这个文章
	if artDao.Id > 0 {
		err = article.UpdateById(c, art)
	} else {
		err = article.Insert(c, art)
	}
	if err != nil {
		return err
	}
	now := time.Now().UnixMilli()
	artDao.UpdatedAt = now
	filter := bson.D{bson.E{Key: "id", Value: artDao.Id},
		bson.E{Key: "user_id", Value: artDao.UserId}}
	set := bson.D{bson.E{Key: "$set", Value: artDao},
		bson.E{Key: "$setOnInsert",
			Value: bson.D{bson.E{Key: "created_at", Value: now}}}}
	_, err = article.pubCol.UpdateOne(c,
		filter, set,
		options.Update().SetUpsert(true))
	return err
}

func (article *ArticleMongoDB) Withdraw(c context.Context, uid int64, id int64, status uint8) (err error) {
	filter := bson.D{bson.E{Key: "id", Value: id},
		bson.E{Key: "user_id", Value: uid}}
	set := bson.D{bson.E{Key: "$set",
		Value: bson.D{bson.E{Key: "status", Value: status}}}}
	res, err := article.col.UpdateOne(c, filter, set)
	if err != nil {
		return err
	}
	if res.ModifiedCount != 1 {
		zap.L().Error("创作者ID不正确")
		return errors.New("你不是作者（ID不正确）")
	}
	_, err = article.pubCol.UpdateOne(c, filter, set)
	return err
}

func (article *ArticleMongoDB) GetById(ctx context.Context, id int64) (art entity.Article, err error) {
	filter := bson.M{"id": id}
	var res artil.Article
	err = article.pubCol.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return entity.Article{}, err
	}
	art = convertor.ToEntity(res)
	return art, nil
}

func (article *ArticleMongoDB) GetByAuthor(ctx context.Context, id int64, offset int64, limit int64) (artList []entity.Article, err error) {
	filter := bson.M{"user_id": id}
	List, err := article.pubCol.Find(ctx, filter, options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer List.Close(ctx)
	for List.Next(ctx) {
		var art artil.Article
		if err = List.Decode(&art); err != nil {
			return nil, err
		}
		artList = append(artList, convertor.ToEntity(art))
	}
	if err = List.Err(); err != nil {
		return nil, err
	}
	return artList, nil
}

func (article *ArticleMongoDB) GetPubById(ctx context.Context, id int64) (art entity.Article, err error) {
	filter := bson.M{"id": id}
	var res artil.Article
	err = article.pubCol.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return entity.Article{}, err
	}
	art = convertor.ToEntity(res)
	return art, nil
}

func (article *ArticleMongoDB) GetList(ctx context.Context, startTime int64, offset int64, limit int64) (artList []entity.Article, err error) {
	const ArticleStatusPublished = 2
	filter := bson.D{
		bson.E{Key: "updated_at", Value: bson.M{"$lt": startTime}},
		bson.E{Key: "status", Value: ArticleStatusPublished},
	}
	List, err := article.pubCol.Find(ctx, filter, options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer List.Close(ctx)
	for List.Next(ctx) {
		var art artil.Article
		if err = List.Decode(&art); err != nil {
			return nil, err
		}
		artList = append(artList, convertor.ToEntity(art))
	}
	if err = List.Err(); err != nil {
		return nil, err
	}
	return artList, nil
}
