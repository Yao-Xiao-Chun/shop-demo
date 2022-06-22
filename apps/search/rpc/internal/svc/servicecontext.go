package svc

import (
	"github.com/meilisearch/meilisearch-go"
	"shop-demo/apps/search/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	SearchServe meilisearch.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	//配置搜索引擎地址
	searchClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   c.MeiliSearch.Host,
		APIKey: c.MeiliSearch.MasterKey,
	})

	return &ServiceContext{
		Config:      c,
		SearchServe: *searchClient,
	}
}
