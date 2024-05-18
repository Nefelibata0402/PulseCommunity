package dao

import (
	"github.com/olivere/elastic/v7"
	"newsCenter/search/infrastructure/persistence/database/search"
)

type SyncElastic struct {
	client *elastic.Client
}

func NewSyncElasticDao() *SyncElastic {
	return &SyncElastic{
		client: search.New(),
	}
}
