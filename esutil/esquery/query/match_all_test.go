package query

import (
	. "github.com/kubemeta/cle-helper/esutil/esquery/common"
	"testing"
)

func TestMatchAll(t *testing.T) {
	RunMapTests(t, []MapTest{
		{
			"match_all without a boost",
			MatchAll(),
			map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
		},
		{
			"match_all with a boost",
			MatchAll().Boost(2.3),
			map[string]interface{}{
				"match_all": map[string]interface{}{
					"boost": 2.3,
				},
			},
		},
		{
			"match_none",
			MatchNone(),
			map[string]interface{}{
				"match_none": map[string]interface{}{},
			},
		},
	})
}
