package esapi

import (
	"context"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

func CreatePolicy(ctx context.Context, es *elasticsearch.Client, policy string, obj io.Reader) (*esapi.Response, error) {
	body := esutil.NewJSONReader(obj)
	req := esapi.ILMPutLifecycleRequest{
		Policy: policy,
		Body:   body,
	}

	return req.Do(ctx, es)
}
