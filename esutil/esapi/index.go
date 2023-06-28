package esapi

import (
	"context"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

const (
	LogSpaceIndex = "system-log-space"
	LogTagIndex   = "system-log-tag"
)

// IndexMapping GET /<target>/_mapping
func IndexMapping(ctx context.Context, es *elasticsearch.Client, index string) (*esapi.Response, error) {
	// Perform the search request.

	req := esapi.IndicesGetMappingRequest{
		Index:  []string{index},
		Pretty: true,
	}

	return req.Do(ctx, es)
}

// ExistIndex HEAD /<target>
func ExistIndex(ctx context.Context, es *elasticsearch.Client, index string) (*esapi.Response, error) {
	req := esapi.CatIndicesRequest{
		Index:  []string{index},
		Format: "json",
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateIndex PUT /<target>
func CreateIndex(ctx context.Context, es *elasticsearch.Client, index string, body io.Reader) (*esapi.Response, error) {
	req := esapi.IndicesCreateRequest{
		Index: index,
	}

	if body != nil {
		req.Body = body
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, err
	}

	return res, nil
}
