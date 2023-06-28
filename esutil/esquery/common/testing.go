package common

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MapTest struct {
	Name string
	M    Mappable
	Exp  map[string]interface{}
}

func RunMapTests(t *testing.T, tests []MapTest) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			m := test.M.Map()

			// convert both maps to JSON in order to compare them. we do not
			// use reflect.DeepEqual on the maps as this doesn't always work
			exp, got, ok := sameJSON(test.Exp, m)
			if !ok {
				t.Errorf("expected %s, got %s", exp, got)
			}
		})
	}
}

func sameJSON(a, b map[string]interface{}) (aJSON, bJSON []byte, ok bool) {
	aJSON, aErr := json.Marshal(a)
	bJSON, bErr := json.Marshal(b)

	if aErr != nil || bErr != nil {
		return aJSON, bJSON, false
	}

	ok = reflect.DeepEqual(aJSON, bJSON)
	return aJSON, bJSON, ok
}

type JsonTest struct {
	Name    string
	M       json.Marshaler
	ExpJSON string
	ExpErr  error
}

func RunJSONTests(t *testing.T, tests []JsonTest) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b, err := test.M.MarshalJSON()
			assert.Equal(t, test.ExpErr, err)
			assert.Equal(t, test.ExpJSON, string(b))
		})
	}
}
