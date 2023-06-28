package aggs

import (
	"github.com/kubemeta/cle-helper/esutil/esquery/common"
)

type NestedAggregation struct {
	name string
	path string
	aggs []common.Aggregation
}

// NestedAgg creates a new aggregation of type "nested". The method name includes
// the "Agg" suffix to prevent conflict with the "nested" query.
func NestedAgg(name string, path string) *NestedAggregation {
	return &NestedAggregation{
		name: name,
		path: path,
	}
}

// Name returns the name of the aggregation.
func (agg *NestedAggregation) Name() string {
	return agg.name
}

// Path NumberOfFragments sets the aggregations path
func (agg *NestedAggregation) Path(p string) *NestedAggregation {
	agg.path = p
	return agg
}

// Aggs sets sub-aggregations for the aggregation.
func (agg *NestedAggregation) Aggs(aggs ...common.Aggregation) *NestedAggregation {
	agg.aggs = aggs
	return agg
}

func (agg *NestedAggregation) Map() map[string]interface{} {
	innerMap := map[string]interface{}{
		"path": agg.path,
	}

	outerMap := map[string]interface{}{
		"nested": innerMap,
	}

	if len(agg.aggs) > 0 {
		subAggs := make(map[string]map[string]interface{})
		for _, sub := range agg.aggs {
			subAggs[sub.Name()] = sub.Map()
		}
		outerMap["aggs"] = subAggs
	}

	return outerMap
}
