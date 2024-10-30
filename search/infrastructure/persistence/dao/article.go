package dao

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"pulseCommunity/search/domain/entity"
	"pulseCommunity/search/infrastructure/persistence/convertor"
	"pulseCommunity/search/infrastructure/persistence/database/search"
	"strconv"
	"strings"
)

const ArticleIndexName = "article_index"

func (s *SyncElastic) InputArticle(ctx context.Context, article entity.Article) error {
	_, err := s.client.Index().
		Index(ArticleIndexName).
		Id(strconv.FormatInt(article.Id, 10)).
		BodyJson(article).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *SyncElastic) SearchArticle(ctx context.Context, keywords []string) ([]entity.Article, error) {
	queryString := strings.Join(keywords, " ")
	status := elastic.NewTermQuery("Status", 2)
	title := elastic.NewMatchQuery("Title", queryString)
	content := elastic.NewMatchQuery("Content", queryString)
	or := elastic.NewBoolQuery().Should(title, content)
	query := elastic.NewBoolQuery().Must(status, or)
	resp, err := s.client.Search(ArticleIndexName).Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]entity.Article, 0, len(resp.Hits.Hits))
	for _, hit := range resp.Hits.Hits {
		var art search.Article
		err = json.Unmarshal(hit.Source, &art)
		if err != nil {
			return nil, err
		}
		article := convertor.ToArticleEntity(art)
		res = append(res, article)
	}
	return res, nil
}
