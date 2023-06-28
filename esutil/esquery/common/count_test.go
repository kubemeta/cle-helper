package common

import (
	. "github.com/kubemeta/cle-helper/esutil/esquery/query"
	"testing"
)

func TestCount(t *testing.T) {
	RunMapTests(t, []MapTest{
		{
			"a simple count request",
			Count(MatchAll()),
			map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{},
				},
			},
		},
	})
}
