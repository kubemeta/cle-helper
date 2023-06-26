package aggs

import (
	"testing"

	. "github.com/kubemeta/cle-helper/esquery/common"
)

func TestNestedAggs(t *testing.T) {
	RunMapTests(t, []MapTest{
		{
			"nested agg: simple",
			NestedAgg("simple", "categories"),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "categories",
				},
			},
		},
		{
			"nested agg: with aggs",
			NestedAgg("more_nested", "authors").
				Aggs(TermsAgg("authors", "name")),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "authors",
				},
				"aggs": map[string]interface{}{
					"authors": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "name",
						},
					},
				},
			},
		},
	})
}
