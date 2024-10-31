package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"pulseCommunity/common/snowflake"
	"testing"
	"time"
)

// Article 数据结构
type Article struct {
	Id        uint64 `json:"id" bson:"id,omitempty"`
	CreatedAt int64  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt int64  `json:"deleted_at" bson:"deleted_at,omitempty"`
	UserId    uint64 `json:"user_id" bson:"user_id,omitempty"`
	Content   string `json:"content" bson:"content,omitempty"`
	Category  string `json:"category" bson:"category"`
	Title     string `json:"title" bson:"title"`
	Status    uint8  `bson:"status,omitempty"`
}

type ArticleEs struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Status  int32  `json:"status"`
	Content string `json:"content"`
}

type Interactive struct {
	Id         int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	BizId      int64  `json:"biz_id" gorm:"uniqueIndex:biz_type_id"`
	Biz        string `json:"biz" gorm:"type:varchar(128);uniqueIndex:biz_type_id"`
	ReadCnt    int64  `json:"read_cnt"`
	LikeCnt    int64  `json:"like_cnt"`
	CollectCnt int64  `json:"collect_cnt"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

//func (*Interactive) TableName() string {
//	return "interactive"
//}

//const ArticleIndexName = "article_index"
//
//func (s *SyncElastic) InputArticle(ctx context.Context, article entity.Article) error {
//	_, err := s.client.Index().
//		Index(ArticleIndexName).
//		Id(strconv.FormatInt(article.Id, 10)).
//		BodyJson(article).Do(ctx)
//	if err != nil {
//		return err
//	}
//	return nil
//}

// 初始化 MongoDB 客户端
func initMongoDB(uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func TestInsertArticles(t *testing.T) {
	// MongoDB 连接 URI 和数据库集合
	uri := "mongodb://root:123456@localhost:27017" // 替换为实际 MongoDB 地址
	client, err := initMongoDB(uri)
	assert.NoError(t, err)
	defer client.Disconnect(context.TODO())

	collection := client.Database("newsCenter").Collection("articles")
	collection1 := client.Database("newsCenter").Collection("published_articles")

	// 定义内容类别
	contentCategories := []string{"Golang社区", "PHP社区", "Python社区", "Java社区", "仓颉社区", "C/C++社区"}
	snowflake.Init(1)
	usedIDs := make(map[uint64]bool) // 用于存储生成过的 ID
	// 批量插入10000条数据
	var articles []interface{}
	for i := 0; i < 10; i++ {
		var artId uint64
		for {
			// 生成唯一的 artId
			artId, err = snowflake.GetID()
			assert.NoError(t, err)
			if !usedIDs[artId] {
				usedIDs[artId] = true // 标记 ID 已使用
				break
			}
		}
		assert.NoError(t, err)
		now := time.Now().UnixMilli()
		article := Article{
			Id:        artId,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: 0,
			UserId:    8279823273164801,
			Content:   "test数据",
			Category:  contentCategories[rand.Intn(len(contentCategories))], // 随机选择一个 Category
			Title:     "test标题",
			Status:    2,
		}
		articles = append(articles, article)
	}

	// 执行批量插入
	insertManyResult, err := collection.InsertMany(context.TODO(), articles)
	insertManyResult1, err := collection1.InsertMany(context.TODO(), articles)
	assert.NoError(t, err)
	fmt.Printf("Inserted %d documents\n", len(insertManyResult.InsertedIDs))
	fmt.Printf("Inserted %d documents\n", len(insertManyResult1.InsertedIDs))
}
