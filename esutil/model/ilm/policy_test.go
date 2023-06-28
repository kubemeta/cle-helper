package ilm

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIlmModel(t *testing.T) {
	hot := NewPhaseSetting(
		//WithPriority(10),
		//WithUnfollow(),
		//WithRollover(1, "day", 50, "GB"),
		//WithReadOnly(),
		//WithShrink(1, 50, "gb"),
		//WithAllocate(2, 200, "include", CustomAttribute{
		//	"box_type": "hot,warm",
		//}),
		WithDelete(nil),
	)
	data, err := json.Marshal(hot)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(data))
}
