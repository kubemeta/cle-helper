package index

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/kubemeta/cle-helper/esutil/model/datastream"
)

type IndexTemplate struct {
	IndexPatterns []string               `json:"index_patterns,omitempty"`
	DataStream    *datastream.DataStream `json:"data_stream,omitempty"`
	ComposedOf    []string               `json:"composed_of,omitempty"`
}

type IndexTemplateFunc func(template *IndexTemplate)

func (t IndexTemplate) Marshal() ([]byte, error) {
	return jsoniter.Marshal(t)
}

func NewIndexTemplate(templateFunc ...IndexTemplateFunc) *IndexTemplate {
	template := &IndexTemplate{}
	for _, f := range templateFunc {
		f(template)
	}

	return template
}

func WithIndexPatterns(patterns ...string) IndexTemplateFunc {
	return func(template *IndexTemplate) {
		template.IndexPatterns = patterns
	}
}

func WithDataStream(stream *datastream.DataStream) IndexTemplateFunc {
	return func(template *IndexTemplate) {
		template.DataStream = stream
	}
}

func WithComposedOf(composed ...string) IndexTemplateFunc {
	return func(template *IndexTemplate) {
		template.ComposedOf = composed
	}
}
