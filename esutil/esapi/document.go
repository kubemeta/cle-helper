package esapi

import (
	"context"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/pkg/errors"
)

// CreateDocument POST /<target>/_doc/ 自动创建文档ID
func CreateDocument(ctx context.Context, es *elasticsearch.Client, index string, doc io.Reader) (*esapi.Response, error) {
	req := esapi.IndexRequest{
		Index:   index,
		Body:    doc,
		Refresh: "true",
	}
	// Perform the request with the client.
	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting response")
	}
	return res, nil
}

// BulkDocument POST /<target>/_doc/bulk 自动创建文档ID
func BulkDocument(ctx context.Context, es *elasticsearch.Client, index string, doc io.Reader) (*esapi.Response, error) {
	req := esapi.BulkRequest{
		Index:   index,
		Body:    doc,
		Refresh: "true",
	}
	// Perform the request with the client.
	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting response")
	}
	return res, nil
}

// GetDocument GET /<target>/_doc/<_id>
func GetDocument(ctx context.Context, es *elasticsearch.Client, index, docId string) (*esapi.Response, error) {
	req := esapi.GetRequest{
		Index:      index,
		DocumentID: docId,
	}
	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting response")
	}
	return res, err
}

// CheckDocumentExistWithQuery GET /<target>/_search query
func CheckDocumentExistWithQuery(ctx context.Context, es *elasticsearch.Client, index string, query io.Reader) (*esapi.Response, error) {
	return SearchExist(ctx, es, index, query)
}

// DeleteDocument DELETE /<target>/_doc/<_id>
func DeleteDocument(ctx context.Context, es *elasticsearch.Client, index, docId string) (*esapi.Response, error) {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: docId,
	}
	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting response")
	}
	return res, err
}

// GetSource GET /<target>/_source/<_id>
func GetSource(ctx context.Context, es *elasticsearch.Client, index, docId string) (*esapi.Response, error) {
	req := esapi.ExistsSourceRequest{
		Index:      index,
		DocumentID: docId,
	}
	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting response")
	}
	return res, err
}

// UpdateDocument POST /<target>/_update/<_id>
func UpdateDocument(ctx context.Context, es *elasticsearch.Client, index string, doc io.Reader, docId string) (*esapi.Response, error) {
	req := esapi.UpdateRequest{
		Index:      index,
		Body:       doc,
		DocumentID: docId,
		Refresh:    "true",
	}
	// Perform the request with the client.
	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting response")
	}

	return res, nil
}

type UpdateQuery struct {
	Doc        interface{} `json:"doc"`
	DetectNoop bool        `json:"detect_noop"`
}

func GenDocumentUpdateQuery(v interface{}) *UpdateQuery {
	return &UpdateQuery{
		Doc:        v,
		DetectNoop: false,
	}
}
