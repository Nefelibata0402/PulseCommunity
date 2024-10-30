package convertor

import (
	"pulseCommunity/article/domain/entity"
	"pulseCommunity/article/infrastructure/persistence/database/article"
	"time"
)

func ToDao(art entity.Article) article.Article {
	return article.Article{
		Id:       art.Id,
		UserId:   art.Author.Id,
		Content:  art.Content,
		Category: art.Category,
		Title:    art.Title,
		Status:   art.Status,
	}
}

func ToEntity(res article.Article) entity.Article {
	return entity.Article{
		Id:        res.Id,
		CreatedAt: time.UnixMilli(res.CreatedAt),
		UpdatedAt: time.UnixMilli(res.UpdatedAt),
		DeletedAt: time.UnixMilli(res.DeletedAt),
		Author: entity.Author{
			Id: res.UserId,
		},
		Content:  res.Content,
		Category: res.Category,
		Title:    res.Title,
		Status:   res.Status,
	}
}
