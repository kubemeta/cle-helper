package common

import (
	"testing"

	. "github.com/kubemeta/cle-helper/esquery/query"
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
