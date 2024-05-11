package articleModel

import "github.com/go-playground/validator/v10"

type ArticleRequest struct {
	ArticleId uint64 `json:"article_id"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Category  string `json:"category"`
}

func ValidateArticleRequest(articleReq *ArticleRequest) error {
	validate := validator.New()
	return validate.Struct(articleReq)
}

type ArticleWithdrawRequest struct {
	ArticleId uint64 `json:"article_id" validate:"required"`
}

func ValidateArticleWithdrawRequest(articleWithdrawReq *ArticleWithdrawRequest) error {
	validate := validator.New()
	return validate.Struct(articleWithdrawReq)
}

type Page struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type Like struct {
	Id   int64 `json:"id"`
	Like bool  `json:"like"`
}

func ValidateLikeRequest(likeReq *Like) error {
	validate := validator.New()
	return validate.Struct(likeReq)
}
