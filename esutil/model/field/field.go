package field

type FieldObject struct {
	Type       string `json:"type"`
	Parameters `json:",inline,omitempty"`
}

type Parameters struct {
	// Only text fields support the analyzer mapping parameter.
	Analyzer            string `json:"analyzer,omitempty"`
	SearchAnalyzer      string `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer string `json:"search_quote_analyzer,omitempty"`
	Boost               int    `json:"boost,omitempty"`
	Coerce              bool   `json:"coerce,omitempty"`
	CopyTo              string `json:"copy_to,omitempty"`
	DocValues           *bool  `json:"doc_values,omitempty"`
	// true, new fields are added to the mapping (default).
	// false, new fields are ignored. These fields will not be indexed or searchable, but will still appear in the _source field of returned hits.
	// runtime, new fields are added to the mapping as runtime fields. These fields are not indexed, and are loaded from _source at query time.
	// strict, if new fields are detected, an exception is thrown and the document is rejected. New fields must be explicitly added to the mapping.
	Dynamic             interface{} `json:"dynamic,omitempty"`
	EagerGlobalOrdinals *bool       `json:"eager_global_ordinals,omitempty"`
	// just store the field without indexing it.
	Enabled *bool  `json:"enabled,omitempty"`
	Format  string `json:"format,omitempty"`
	// Strings longer than the ignore_above setting will not be indexed or stored.
	IgnoreAbove     int   `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool `json:"ignore_malformed,omitempty"`
	Index           *bool `json:"index,omitempty"`
	// docs, only the doc number is indexed
	// freqs, doc number and term frequencies are indexed
	// positions(default), doc number, term frequencies, and term positions (or order) are indexed.
	// offsets, doc number, term frequencies, positions, and start and end character offsets are indexed.
	IndexOptions         string                 `json:"index_options,omitempty"`
	IndexPhrases         *bool                  `json:"index_phrases,omitempty"`
	IndexPrefixes        *IndexPrefixes         `json:"index_prefixes,omitempty"`
	Meta                 *Meta                  `json:"meta,omitempty"`
	Fields               map[string]FieldObject `json:"fields,omitempty"`
	Normalizer           string                 `json:"normalizer,omitempty"`
	Norms                *bool                  `json:"norms,omitempty"`
	NullValue            string                 `json:"null_value,omitempty"`
	PositionIncrementGap int                    `json:"position_increment_gap,omitempty"`
	// BM25
	// classic
	// boolean
	Similarity string `json:"similarity,omitempty"`
	Store      *bool  `json:"store,omitempty"`
	// the term_vector setting accepts:
	// no,yes, with_positions, with_offsets, with_positions_offsets, with_positions_payloads, with_positions_offsets_payloads
	TermVector string `json:"term_vector,omitempty"`
}

type IndexPrefixes struct {
	MinChars int `json:"min_chars,omitempty"`
	MaxChars int `json:"max_chars,omitempty"`
}

type Meta struct {
	Unit       string `json:"unit,omitempty"`
	MetricType string `json:"metricType,omitempty"`
}

const (
	Binary  = "binary"
	Boolean = "boolean"
)

type Keywords string

const (
	Keyword         = Keywords("keyword")
	ConstantKeyword = Keywords("constant_keyword")
	Wildcard        = Keywords("wildcard")
)

type Numeric string

const (
	// Long is a signed 64-bit integer with a minimum value of -263 and a maximum value of 263-1.
	Long = Numeric("long")
	// Integer is a signed 32-bit integer with a minimum value of -231 and a maximum value of 231-1.
	Integer = Numeric("integer")
	// Short is a signed 16-bit integer with a minimum value of -32,768 and a maximum value of 32,767.
	Short = Numeric("short")
	// Byte is a signed 8-bit integer with a minimum value of -128 and a maximum value of 127.
	Byte = Numeric("byte")
	// Double is a double-precision 64-bit IEEE 754 floating point number, restricted to finite values.
	Double = Numeric("double")
	// Float a single-precision 32-bit IEEE 754 floating point number, restricted to finite values.
	Float = Numeric("float")
	// HalfFloat a half-precision 16-bit IEEE 754 floating point number, restricted to finite values.
	HalfFloat = Numeric("half_float")
	// ScaledFloat a floating point number that is backed by a long, scaled by a fixed double scaling factor.
	ScaledFloat = Numeric("scaled_float")
	// UnsignedLong an unsigned 64-bit integer with a minimum value of 0 and a maximum value of 264-1.
	UnsignedLong = Numeric("unsigned_long")
)

type Dates string

const (
	Date      = Dates("date")
	DateNanos = Dates("date_nanos")
)

/*
Alias
1. The target must be a concrete field, and not an object or another field alias.
2. The target field must exist at the time the alias is created.
3. If nested objects are defined, a field alias must have the same nested scope as its target.
4. A field alias can only have one target.
*/

const (
	Alias = "alias"
)

type AliasObject struct {
	Type string `json:"type"`
	// The path to the target field. Note that this must be the full path, including any parent objects.(object1.object2.field)
	Path string `json:"path"`
}

type Objects string

const (
	Object    = Objects("object")
	Flattened = Objects("flattened")
	Nested    = Objects("nested")
	Join      = Objects("join")
)

type Structured string

const (
	IntegerRange = Structured("integer_range")
	FloatRange   = Structured("float_range")
	LongRange    = Structured("long_range")
	DoubleRange  = Structured("double_range")
	DateRange    = Structured("date_range")
	IpRange      = Structured("ip_range")
	IP           = Structured("ip")
	Version      = Structured("version")
	Murmur3      = Structured("murmur3")
)

type Aggregate string

const (
	AggregateMetricDouble = Aggregate("aggregate_metric_double")
	Histogram             = Aggregate("histogram")
)

type TextSearch string

const (
	Text = TextSearch("text")
)

type Spatial string

const (
	GeoPoint = Spatial("geo_point")
	GeoShape = Spatial("geo_shape")
	Point    = Spatial("point")
	Shape    = Spatial("shape")
)
