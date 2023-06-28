package aggs

import (
	"github.com/kubemeta/cle-helper/esutil/esquery/common"
)

type FilterAggregation struct {
	name   string
	filter common.Mappable
	aggs   []common.Aggregation
}

// FilterAgg creates a new common.Aggregation of type "filter". The method name includes
// the "Agg" suffix to prevent conflict with the "filter" query.
func FilterAgg(name string, filter common.Mappable) *FilterAggregation {
	return &FilterAggregation{
		name:   name,
		filter: filter,
	}
}

// Name returns the name of the common.Aggregation.
func (agg *FilterAggregation) Name() string {
	return agg.name
}

// Filter sets the filter items
func (agg *FilterAggregation) Filter(filter common.Mappable) *FilterAggregation {
	agg.filter = filter
	return agg
}

// Aggs sets sub-common.Aggregations for the common.Aggregation.
func (agg *FilterAggregation) Aggs(aggs ...common.Aggregation) *FilterAggregation {
	agg.aggs = aggs
	return agg
}

func (agg *FilterAggregation) Map() map[string]interface{} {
	outerMap := map[string]interface{}{
		"filter": agg.filter.Map(),
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
