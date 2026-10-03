// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ptraceotlp

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

var _ json.Unmarshaler = ExportRequest{}
var _ json.Marshaler = ExportRequest{}

var tracesRequestJSON = []byte(`
	{
		"resourceSpans": [
			{
				"resource": {},
				"scopeSpans": [
					{
						"scope": {},
						"spans": [
							{
								"traceId": "00000000000000000000000000000000",
								"spanId": "0000000000000000",
								"parentSpanId": "0000000000000000",
								"name": "test_span",
								"status": {}
							}
						]
					}
				]
			}
		]
	}`)

func TestRequestToPData(t *testing.T) {
	tr := NewExportRequest()
	assert.Equal(t, tr.Traces().SpanCount(), 0)
	tr.Traces().ResourceSpans().AppendEmpty().ScopeSpans().AppendEmpty().Spans().AppendEmpty()
	assert.Equal(t, tr.Traces().SpanCount(), 1)
}

func TestRequestJSON(t *testing.T) {
	tr := NewExportRequest()
	assert.NoError(t, tr.UnmarshalJSON(tracesRequestJSON))
	assert.Equal(t, "test_span", tr.Traces().ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).Name())

	got, err := tr.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, strings.Join(strings.Fields(string(tracesRequestJSON)), ""), string(got))
}

// An attribute with a key and no value is valid protobuf. Both decoders
// must give a request that reads the attribute as an empty value.
func TestRequestUnmarshalAttributeWithoutValue(t *testing.T) {
	// ExportTraceServiceRequest{resource_spans: [{resource: {attributes:
	// [{key: "k"}]}}]}, with no value field in the KeyValue.
	body := []byte{0x0a, 0x07, 0x0a, 0x05, 0x0a, 0x03, 0x0a, 0x01, 'k'}
	for name, unmarshal := range map[string]func(ExportRequest, []byte) error{
		"copy":      ExportRequest.UnmarshalProto,
		"zero-copy": ExportRequest.UnmarshalProtoUnsafe,
	} {
		t.Run(name, func(t *testing.T) {
			req := NewExportRequest()
			require.NoError(t, unmarshal(req, append([]byte(nil), body...)))
			attrs := req.Traces().ResourceSpans().At(0).Resource().Attributes()
			v, ok := attrs.Get("k")
			require.True(t, ok)
			assert.Equal(t, pcommon.ValueTypeEmpty, v.Type())
			assert.Equal(t, "", v.AsString())
			attrs.PutStr("k", "v")
			assert.Equal(t, map[string]any{"k": "v"}, attrs.AsRaw())
		})
	}
}
