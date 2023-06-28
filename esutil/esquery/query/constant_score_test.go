package query

import (
	. "github.com/kubemeta/cle-helper/esutil/esquery/common"
	"testing"
)

func TestConstantScore(t *testing.T) {
	RunMapTests(t, []MapTest{
		{
			"constant_score query without boost",
			ConstantScore(Term("user", "kimchy")),
			map[string]interface{}{
				"constant_score": map[string]interface{}{
					"filter": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
				},
			},
		},
		{
			"constant_score query with boost",
			ConstantScore(Term("user", "kimchy")).Boost(2.2),
			map[string]interface{}{
				"constant_score": map[string]interface{}{
					"filter": map[string]interface{}{
						"term": map[string]interface{}{
							"user": map[string]interface{}{
								"value": "kimchy",
							},
						},
					},
					"boost": 2.2,
				},
			},
		},
	})
}
