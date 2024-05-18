package repository

import (
	"context"
	"newsCenter/search/domain/entity"
)

type UserRepository interface {
	InputUser(ctx context.Context, article entity.User) error
	SearchUser(ctx context.Context, keywords []string) ([]entity.User, error)
}

type ArticleRepository interface {
	InputArticle(ctx context.Context, article entity.Article) error
	SearchArticle(ctx context.Context, keywords []string) ([]entity.Article, error)
}
