package test

import (
	. "github.com/kubemeta/cle-helper/esutil/esquery/common"
	. "github.com/kubemeta/cle-helper/esutil/esquery/query"
	. "github.com/kubemeta/cle-helper/esutil/esquery/search"
	"testing"
)

func TestQueryMaps(t *testing.T) {
	RunMapTests(t, []MapTest{
		{
			"a simple match_all query",
			Query(MatchAll()),
			map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{},
				},
			},
		},
		{
			"a complex query",
			Query(
				Bool().
					Must(
						Range("date").
							Gt("some time in the past").
							Lte("now").
							Relation(RangeContains).
							TimeZone("Asia/Jerusalem").
							Boost(2.3),

						Match("author").
							Query("some guy").
							Analyzer("analyzer?").
							Fuzziness("fuzz"),
					).
					Boost(3.1),
			),
			map[string]interface{}{
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []map[string]interface{}{
							{
								"range": map[string]interface{}{
									"date": map[string]interface{}{
										"gt":        "some time in the past",
										"lte":       "now",
										"relation":  "CONTAINS",
										"time_zone": "Asia/Jerusalem",
										"boost":     2.3,
									},
								},
							},
							{
								"match": map[string]interface{}{
									"author": map[string]interface{}{
										"query":     "some guy",
										"analyzer":  "analyzer?",
										"fuzziness": "fuzz",
									},
								},
							},
						},
						"boost": 3.1,
					},
				},
			},
		},
	})
}

func TestQueryJSONs(t *testing.T) {
	RunJSONTests(t, []JsonTest{
		{
			"simple query",
			Query(
				Bool().
					Must(Term("account_id", "bla")),
			),
			`{"query":{"bool":{"must":[{"term":{"account_id":{"value":"bla"}}}]}}}`,
			nil,
		},
	})
}
