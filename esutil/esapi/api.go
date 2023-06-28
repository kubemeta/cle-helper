package esapi

import (
	"context"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ApiBuilder interface {
	Create(ctx context.Context, es *elasticsearch.Client, request interface{}) (*esapi.Response, error)
}
