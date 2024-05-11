package entity

import "time"

type Article struct {
	Id        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Author    Author
	Content   string
	Category  string
	Title     string
	Status    uint8
}
type Author struct {
	Id   uint64
	Name string
}

func (art Article) Abstract() string {
	str := []rune(art.Content)
	if len(str) > 128 {
		str = str[:128]
	}
	return string(str)
}

const (
	ArticleStatus = iota
	// ArticleStatusNoPublish 文章为发布
	ArticleStatusNoPublish
	// ArticleStatusPublish 文章发布
	ArticleStatusPublish
	// ArticleStatusWithdraw 文章撤回
	ArticleStatusWithdraw
)

type Interactive struct {
	ReadCnt    int64
	LikeCnt    int64
	CollectCnt int64
	Liked      bool
	Collected  bool
}
