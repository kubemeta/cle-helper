package component

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"

	"github.com/kubemeta/cle-helper/esutil/model/ilm"
	"github.com/kubemeta/cle-helper/esutil/model/unit"
)

type ComponentTemplate struct {
	Template Template `json:"template"`
	Version  string   `json:"version,omitempty"`
}

func (c ComponentTemplate) Marshal() ([]byte, error) {
	return jsoniter.Marshal(c)
}

type Template struct {
	Aliases  Aliases   `json:"aliases,omitempty"`
	Mappings *Mappings `json:"mappings,omitempty"`
	Settings *Settings `json:"settings,omitempty"`
}

type Aliases map[string]interface{}

type Mappings struct {
	DateDetection      *bool  `json:"date_detection,omitempty"`
	DynamicDateFormats string `json:"dynamic_date_formats,omitempty"`
	NumericDetection   *bool  `json:"numeric_detection,omitempty"`
	// Dynamic Mapping
	DynamicTemplates []map[string]DynamicTemplate `json:"dynamic_templates,omitempty"`
	// Explicit Mapping
	Properties Properties `json:"properties,omitempty"`
}

type Properties map[string]interface{}

func (p Properties) Extend(properties Properties) Properties {
	for k, v := range properties {
		p[k] = v
	}
	return p
}

type DynamicTemplate struct {
	// The match conditions
	MatchMappingType string `json:"match_mapping_type,omitempty"`
	Match            string `json:"match,omitempty"`
	MatchPattern     string `json:"match_pattern,omitempty"`
	UnMatch          string `json:"unmatch,omitempty"`
	PatchMatch       string `json:"patch_match,omitempty"`
	PatchUnMatch     string `json:"patch_unmatch,omitempty"`
	// Runtime Field
	Runtime *Runtime `json:"runtime,omitempty"`
	// The mapping that the matched field should use.
	Mapping *Mapping `json:"mapping,omitempty"`
}

type Runtime struct {
	Type string `json:"type,omitempty"`
}

type Mapping map[string]interface{}

type MappingsFunc func(mappings *Mappings)

func NewMappings(mappingsFunc ...MappingsFunc) *Mappings {
	m := &Mappings{}
	for _, mf := range mappingsFunc {
		mf(m)
	}
	return m
}

func WithDateDetection(dateDetection *bool) MappingsFunc {
	return func(mappings *Mappings) {
		mappings.DateDetection = dateDetection
	}
}

func WithDynamicDateFormats(format string) MappingsFunc {
	return func(mappings *Mappings) {
		mappings.DynamicDateFormats = format
	}
}

func WithNumberDetection(numericDetection bool) MappingsFunc {
	return func(mappings *Mappings) {
		mappings.NumericDetection = &numericDetection
	}
}

func WithDynamicTemplates(dynamicTemplates []map[string]DynamicTemplate) MappingsFunc {
	return func(mappings *Mappings) {
		mappings.DynamicTemplates = dynamicTemplates
	}
}

func WithProperties(properties map[string]interface{}) MappingsFunc {
	return func(mappings *Mappings) {
		mappings.Properties = properties
	}
}

type Settings struct {
	// static index settings
	// default to 1
	NumberOfShards                int          `json:"number_of_shards,omitempty"`
	NumberOfRoutingShards         int          `json:"number_of_routing_shards,omitempty"`
	Codec                         string       `json:"codec,omitempty"`
	RoutingPartitionSize          int          `json:"routing_partition_size,omitempty"`
	SoftDeletes                   *SoftDeletes `json:"soft_deletes,omitempty"`
	LoadFixedBitsetFiltersEagerly *bool        `json:"load_fixed_bitset_filters_eagerly,omitempty"`

	// dynamic index settings
	// default to 1
	NumberOfReplicas   int     `json:"number_of_replicas,omitempty"`
	AutoExpandReplicas *bool   `json:"auto_expand_replicas,omitempty"`
	Search             *Search `json:"search,omitempty"`
	// default to 1s
	RefreshInterval      int `json:"refresh_interval,omitempty"`
	MaxResultWindow      int `json:"max_result_window,omitempty"`
	MaxInnerResultWindow int `json:"max_inner_result_window,omitempty"`
	MaxReScoreWindow     int `json:"max_rescore_window,omitempty"`
	// default to 100
	MaxDocValueFieldsSearch int `json:"max_docvalue_fields_search,omitempty"`
	// default to 32
	MaxScriptFields int `json:"max_script_fields,omitempty"`
	// default to 1
	MaxNgrmDiff int `json:"max_ngrm_diff,omitempty"`
	// default to 3
	MaxShingleDiff      int        `json:"max_shingle_diff,omitempty"`
	MaxRefreshListeners int        `json:"max_refresh_listeners,omitempty"`
	Analyze             *Analyze   `json:"analyze,omitempty"`
	Highlight           *Highlight `json:"highlight,omitempty"`
	// default to 65536
	MaxTermsCount  int      `json:"max_terms_count,omitempty"`
	MaxRegexLength int      `json:"max_regex_length,omitempty"`
	Query          *Query   `json:"query,omitempty"`
	Routing        *Routing `json:"routing,omitempty"`
	// default to 60s
	GcDeletes       string                    `json:"gc_deletes,omitempty"`
	DefaultPipeline string                    `json:"default_pipeline,omitempty"`
	FinalPipeline   string                    `json:"final_pipeline,omitempty"`
	Hidden          *bool                     `json:"hidden,omitempty"`
	Indices         *ilm.ClusterLevelSettings `json:"indices,omitempty"`
	Index           *ilm.IndexLevelSettings   `json:"index,omitempty"`
}

type SoftDeletes struct {
	RetentionLease *RetentionLease `json:"retention_lease,omitempty"`
}

type RetentionLease struct {
	// default to 12h
	Period string `json:"period"`
}

type Search struct {
	Idle *Idle `json:"idle,omitempty"`
}

type Analyze struct {
	// default to 10000
	MaxTokenCount int `json:"max_token_count,omitempty"`
}

type Highlight struct {
	// default to 1000000
	MaxAnalyzeOffset int `json:"max_analyze_offset,omitempty"`
}

type Idle struct {
	// default to 30s
	After string `json:"after"`
}

type Query struct {
	DefaultField string `json:"default_field,omitempty"`
}

type Routing struct {
}

type SettingsFunc func(settings *Settings)

func NewIndexSetting(settingsFunc ...SettingsFunc) *Settings {
	s := &Settings{}
	for _, f := range settingsFunc {
		f(s)
	}
	return s
}

func WithNumberOfShards(shards int) SettingsFunc {
	return func(settings *Settings) {
		settings.NumberOfShards = shards
	}
}

func WithNumberOfReplicas(replicas int) SettingsFunc {
	return func(settings *Settings) {
		settings.NumberOfReplicas = replicas
	}
}

func WithClusterIndexSettings(historyIndexEnabled bool, pollInterval int, unit unit.TimeUnit) SettingsFunc {
	return func(settings *Settings) {
		settings.Indices = &ilm.ClusterLevelSettings{
			GlobalLifecycle: &ilm.GlobalLifecycle{
				HistoryIndexEnabled: &historyIndexEnabled,
				PollInterval:        fmt.Sprintf("%d%s", pollInterval, unit),
			},
		}
	}
}

func WithIndexSettings(index *ilm.IndexLevelSettings) SettingsFunc {
	return func(settings *Settings) {
		settings.Index = index
	}
}

type AliasesFunc func(aliases Aliases)

func NewAliases(aliasesFunc ...AliasesFunc) Aliases {
	a := Aliases{}
	for _, f := range aliasesFunc {
		f(a)
	}
	return a
}
