package esapi

import (
	"context"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

func CreateIndexTemplate(ctx context.Context, es *elasticsearch.Client, name string, obj interface{}) (*esapi.Response, error) {
	body := esutil.NewJSONReader(obj)
	req := esapi.IndicesPutIndexTemplateRequest{
		Name: name,
		Body: body,
	}

	return req.Do(ctx, es)
}
