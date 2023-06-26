package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"

	. "github.com/kubemeta/cle-helper/esquery/common"
)

// SearchRequest represents a request to ElasticSearch's Search API, described
// in https://www.elastic.co/guide/en/elasticsearch/reference/current/search.html.
// Not all features of the search API are currently supported, but a request can
// currently include a query, aggregations, and more.
type SearchRequest struct {
	aggs        []Aggregation
	explain     *bool
	from        *uint64
	highlight   Mappable
	searchAfter []interface{}
	postFilter  Mappable
	query       Mappable
	size        *uint64
	sort        Sort
	source      Source
	timeout     *time.Duration
}

// Search creates a new SearchRequest object, to be filled via method chaining.
func Search() *SearchRequest {
	return &SearchRequest{}
}

// Query sets a query for the request.
func (req *SearchRequest) Query(q Mappable) *SearchRequest {
	req.query = q
	return req
}

// Aggs sets one or more aggregations for the request.
func (req *SearchRequest) Aggs(aggs ...Aggregation) *SearchRequest {
	req.aggs = append(req.aggs, aggs...)
	return req
}

// PostFilter sets a post_filter for the request.
func (req *SearchRequest) PostFilter(filter Mappable) *SearchRequest {
	req.postFilter = filter
	return req
}

// From sets a document offset to start from.
func (req *SearchRequest) From(offset uint64) *SearchRequest {
	req.from = &offset
	return req
}

// Size sets the number of hits to return. The default - according to the ES
// documentation - is 10.
func (req *SearchRequest) Size(size uint64) *SearchRequest {
	req.size = &size
	return req
}

// Sort sets how the results should be sorted.
func (req *SearchRequest) Sort(name string, order Order) *SearchRequest {
	req.sort = append(req.sort, map[string]interface{}{
		name: map[string]interface{}{
			"order": order,
		},
	})

	return req
}

// SearchAfter retrieve the sorted result
func (req *SearchRequest) SearchAfter(s ...interface{}) *SearchRequest {
	req.searchAfter = append(req.searchAfter, s...)
	return req
}

// Explain sets whether the ElasticSearch API should return an explanation for
// how each hit's score was calculated.
func (req *SearchRequest) Explain(b bool) *SearchRequest {
	req.explain = &b
	return req
}

// Timeout sets a timeout for the request.
func (req *SearchRequest) Timeout(dur time.Duration) *SearchRequest {
	req.timeout = &dur
	return req
}

// SourceIncludes sets the keys to return from the matching documents.
func (req *SearchRequest) SourceIncludes(keys ...string) *SearchRequest {
	req.source.SetIncludes(keys)
	return req
}

// SourceExcludes sets the keys to not return from the matching documents.
func (req *SearchRequest) SourceExcludes(keys ...string) *SearchRequest {
	req.source.SetExcludes(keys)
	return req
}

// Highlight sets a highlight for the request.
func (req *SearchRequest) Highlight(highlight Mappable) *SearchRequest {
	req.highlight = highlight
	return req
}

// Map implements the Mappable interface. It converts the request to into a
// nested map[string]interface{}, as expected by the go-elasticsearch library.
func (req *SearchRequest) Map() map[string]interface{} {
	m := make(map[string]interface{})
	if req.query != nil {
		m["query"] = req.query.Map()
	}
	if len(req.aggs) > 0 {
		aggs := make(map[string]interface{})
		for _, agg := range req.aggs {
			aggs[agg.Name()] = agg.Map()
		}

		m["aggs"] = aggs
	}
	if req.postFilter != nil {
		m["post_filter"] = req.postFilter.Map()
	}
	if req.size != nil {
		m["size"] = *req.size
	}
	if len(req.sort) > 0 {
		m["sort"] = req.sort
	}
	if req.from != nil {
		m["from"] = *req.from
	}
	if req.explain != nil {
		m["explain"] = *req.explain
	}
	if req.timeout != nil {
		m["timeout"] = fmt.Sprintf("%.0fs", req.timeout.Seconds())
	}
	if req.highlight != nil {
		m["highlight"] = req.highlight.Map()
	}
	if req.searchAfter != nil {
		m["search_after"] = req.searchAfter
	}

	source := req.source.Map()
	if len(source) > 0 {
		m["_source"] = source
	}

	return m
}

// MarshalJSON implements the json.Marshaler interface. It returns a JSON
// representation of the map generated by the SearchRequest's Map method.
func (req *SearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(req.Map())
}

// Run executes the request using the provided ElasticSearch client. Zero or
// more search options can be provided as well. It returns the standard Response
// type of the official Go client.
func (req *SearchRequest) Run(
	api *elasticsearch.Client,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	return req.RunSearch(api.Search, o...)
}

// RunSearch is the same as the Run method, except that it accepts a value of
// type esapi.Search (usually this is the Search field of an elasticsearch.Client
// object). Since the ElasticSearch client does not provide an interface type
// for its API (which would allow implementation of mock clients), this provides
// a workaround. The Search function in the ES client is actually a field of a
// function type.
func (req *SearchRequest) RunSearch(
	search esapi.Search,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(req.Map())
	if err != nil {
		return nil, err
	}

	opts := append([]func(*esapi.SearchRequest){search.WithBody(&b)}, o...)

	return search(opts...)
}

// Query is a shortcut for creating a SearchRequest with only a query. It is
// mostly included to maintain the API provided by esquery in early releases.
func Query(q Mappable) *SearchRequest {
	return Search().Query(q)
}

// Aggregate is a shortcut for creating a SearchRequest with aggregations. It is
// mostly included to maintain the API provided by esquery in early releases.
func Aggregate(aggs ...Aggregation) *SearchRequest {
	return Search().Aggs(aggs...)
}
