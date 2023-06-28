package component

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"

	"github.com/kubemeta/cle-helper/esutil/model/ilm"
)

func TestComponentTemplateModel(t *testing.T) {
	m := NewMappings(
		WithDynamicTemplates([]map[string]DynamicTemplate{
			{
				"strings_as_ip": DynamicTemplate{
					MatchMappingType: "string",
					Match:            "ip*",
					UnMatch:          "*v6",
					Runtime: &Runtime{
						Type: "ip",
					},
				},
			},
		}))

	mapping, err := jsoniter.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(mapping))

	s := NewIndexSetting(
		WithNumberOfShards(1),
		WithNumberOfReplicas(1),
		//WithClusterIndexSettings(false, 10, unit.Minutes),
		WithIndexSettings(&ilm.IndexLevelSettings{
			Lifecycle: &ilm.Lifecycle{
				Name: "my_policy",
			},
		}),
	)
	setting, err := jsoniter.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(setting))
}
