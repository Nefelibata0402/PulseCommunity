package convertor

import (
	"newsCenter/search/domain/entity"
	"newsCenter/search/infrastructure/persistence/database/search"
)

//func ToDao(art entity.Article) article.Article {
//	return article.Article{
//		Id:       art.Id,
//		UserId:   art.Author.Id,
//		Content:  art.Content,
//		Category: art.Category,
//		Title:    art.Title,
//		Status:   art.Status,
//	}
//}

func ToUserEntity(res search.User) entity.User {
	return entity.User{
		Id:       res.Id,
		Nickname: res.Nickname,
	}
}

func ToArticleEntity(res search.Article) entity.Article {
	return entity.Article{
		Id:      res.Id,
		Title:   res.Title,
		Status:  res.Status,
		Content: res.Content,
	}
}
