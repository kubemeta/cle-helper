package esapi

import (
	"context"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// Search GET /<target>/_search
func Search(ctx context.Context, es *elasticsearch.Client, index string, filterPath []string, from, size int, query io.Reader) (*esapi.Response, error) {
	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex(index),
		es.Search.WithBody(query),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithFilterPath(filterPath...),
		es.Search.WithFrom(from),
		es.Search.WithSize(size),
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func SearchSets(ctx context.Context, es *elasticsearch.Client, index string, filterPath []string, query io.Reader) (*esapi.Response, error) {
	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex(index),
		es.Search.WithBody(query),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithFilterPath(filterPath...),
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// SearchExist GET /<target>/_search
func SearchExist(ctx context.Context, es *elasticsearch.Client, index string, reader io.Reader) (*esapi.Response, error) {
	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex(index),
		es.Search.WithBody(reader),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// SearchRaw GET /<target>/_search
func SearchRaw(ctx context.Context, es *elasticsearch.Client, index string, filterPath []string, from, size int, query io.Reader) (*esapi.Response, error) {
	// Perform the search request.
	req := esapi.SearchRequest{
		Index:          []string{index},
		Body:           query,
		TrackTotalHits: true,
		Pretty:         true,
		FilterPath:     filterPath,
		From:           &from,
		Size:           &size,
	}

	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, err
	}

	return res, nil
}
