package dao

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"newsCenter/search/domain/entity"
	"newsCenter/search/infrastructure/persistence/convertor"
	"newsCenter/search/infrastructure/persistence/database/search"
	"strconv"
	"strings"
)

const UserIndexName = "user_index"

func (s *SyncElastic) InputUser(ctx context.Context, user entity.User) (err error) {
	_, err = s.client.Index().
		Index(UserIndexName).
		Id(strconv.FormatInt(user.Id, 10)).
		BodyJson(user).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *SyncElastic) SearchUser(ctx context.Context, keywords []string) ([]entity.User, error) {
	queryString := strings.Join(keywords, "")
	query := elastic.NewMatchQuery("username", queryString)
	resp, err := s.client.Search(UserIndexName).
		Query(query).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]entity.User, 0, len(resp.Hits.Hits))
	for _, hit := range resp.Hits.Hits {
		var u search.User
		err = json.Unmarshal(hit.Source, &u)
		if err != nil {
			return nil, err
		}
		user := convertor.ToUserEntity(u)
		res = append(res, user)
	}
	return res, err
}
