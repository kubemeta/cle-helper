package test

import (
	. "github.com/kubemeta/cle-helper/esutil/esquery/common"
	. "github.com/kubemeta/cle-helper/esutil/esquery/search"
	"testing"
)

func TestCustomQuery(t *testing.T) {
	m := map[string]interface{}{
		"geo_distance": map[string]interface{}{
			"distance": "200km",
			"pin.location": map[string]interface{}{
				"lat": 40,
				"lon": -70,
			},
		},
	}

	RunMapTests(t, []MapTest{
		{
			"custom query",
			CustomQuery(m),
			m,
		},
	})
}

func TestCustomAgg(t *testing.T) {
	m := map[string]interface{}{
		"genres": map[string]interface{}{
			"terms": map[string]interface{}{
				"field": "genre",
			},
			"t_shirts": map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"type": "t-shirt",
					},
				},
				"aggs": map[string]interface{}{
					"avg_price": map[string]interface{}{
						"avg": map[string]interface{}{
							"field": "price",
						},
					},
				},
			},
		},
	}

	RunMapTests(t, []MapTest{
		{
			"custom aggregation",
			CustomAgg("custom_agg", m),
			m,
		},
	})
}
