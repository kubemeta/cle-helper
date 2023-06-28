package client

import (
	"github.com/elastic/go-elasticsearch/v7"

	"github.com/kubemeta/cle-helper/esutil/esapi"
)

type EsClient interface {
	Client() (*elasticsearch.Client, error)
}

type esClient struct {
	conf elasticsearch.Config
}

func New(conf elasticsearch.Config) EsClient {
	return &esClient{conf}
}

func (es esClient) Client() (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(es.conf)
	if err != nil {
		return nil, err
	}
	//Ping 的方式检查是否可用
	res, err := client.Ping()
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, esapi.ParseResError(res)
	}
	return client, nil
}
