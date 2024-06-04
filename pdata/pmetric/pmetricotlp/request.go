// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetricotlp // import "github.com/oodle-ai/opentelemetry-collector/pdata/pmetric/pmetricotlp"

import (
	"bytes"

	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlpcollectormetrics "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/collector/metrics/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal/json"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"
)

var jsonUnmarshaler = &pmetric.JSONUnmarshaler{}

// ExportRequest represents the request for gRPC/HTTP client/server.
// It's a wrapper for pmetric.Metrics data.
type ExportRequest struct {
	orig  *otlpcollectormetrics.ExportMetricsServiceRequest
	state *internal.State
}

// NewExportRequest returns an empty ExportRequest.
func NewExportRequest() ExportRequest {
	state := internal.StateMutable
	return ExportRequest{
		orig:  otlpcollectormetrics.ExportMetricsServiceRequestFromVTPool(),
		state: &state,
	}
}

// NewExportRequestFromMetrics returns a ExportRequest from pmetric.Metrics.
// Because ExportRequest is a wrapper for pmetric.Metrics,
// any changes to the provided Metrics struct will be reflected in the ExportRequest and vice versa.
func NewExportRequestFromMetrics(md pmetric.Metrics) ExportRequest {
	return ExportRequest{
		orig:  internal.GetOrigMetrics(internal.Metrics(md)),
		state: internal.GetMetricsState(internal.Metrics(md)),
	}
}

// MarshalProto marshals ExportRequest into proto bytes.
func (ms ExportRequest) MarshalProto(buf *bytes.Buffer) error {
	size := ms.orig.SizeVT()
	buf.Grow(size)
	availableBuf := buf.AvailableBuffer()
	availableBuf = availableBuf[:size]
	_, err := ms.orig.MarshalToVT(availableBuf)
	if err != nil {
		return err
	}

	_, err = buf.Write(availableBuf)
	return err
}

// UnmarshalProto unmarshalls ExportRequest from proto bytes.
func (ms ExportRequest) UnmarshalProto(data []byte) error {
	return ms.orig.UnmarshalVT(data)
}

// UnmarshalProtoUnsafe unmarshalls ExportRequest from proto bytes.
func (ms ExportRequest) UnmarshalProtoUnsafe(data []byte) error {
	return ms.orig.UnmarshalVTUnsafe(data)
}

func (ms ExportRequest) ReturnToPool() {
	ms.orig.ReturnToVTPool()
}

// MarshalJSON marshals ExportRequest into JSON bytes.
func (ms ExportRequest) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if err := json.Marshal(&buf, ms.orig); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON unmarshalls ExportRequest from JSON bytes.
func (ms ExportRequest) UnmarshalJSON(data []byte) error {
	md, err := jsonUnmarshaler.UnmarshalMetrics(data)
	if err != nil {
		return err
	}
	*ms.orig = *internal.GetOrigMetrics(internal.Metrics(md))
	return nil
}

func (ms ExportRequest) Metrics() pmetric.Metrics {
	return pmetric.Metrics(internal.NewMetrics(ms.orig, ms.state))
}
