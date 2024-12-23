// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ptraceotlp // import "github.com/oodle-ai/opentelemetry-collector/pdata/ptrace/ptraceotlp"

import (
	"bytes"

	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlpcollectortrace "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/collector/trace/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal/json"
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal/otlp"
	"github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"
)

var jsonUnmarshaler = &ptrace.JSONUnmarshaler{}

// ExportRequest represents the request for gRPC/HTTP client/server.
// It's a wrapper for ptrace.Traces data.
type ExportRequest struct {
	orig  *otlpcollectortrace.ExportTraceServiceRequest
	state *internal.State
}

// NewExportRequest returns an empty ExportRequest.
func NewExportRequest() ExportRequest {
	state := internal.StateMutable
	return ExportRequest{
		orig:  otlpcollectortrace.ExportTraceServiceRequestFromVTPool(),
		state: &state,
	}
}

// NewExportRequestFromTraces returns a ExportRequest from ptrace.Traces.
// Because ExportRequest is a wrapper for ptrace.Traces,
// any changes to the provided Traces struct will be reflected in the ExportRequest and vice versa.
func NewExportRequestFromTraces(td ptrace.Traces) ExportRequest {
	return ExportRequest{
		orig:  internal.GetOrigTraces(internal.Traces(td)),
		state: internal.GetTracesState(internal.Traces(td)),
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

// UnmarshalProtoUnsafe unmarshalls ExportRequest from proto bytes.
func (ms ExportRequest) UnmarshalProtoUnsafe(data []byte) error {
	return ms.orig.UnmarshalVTUnsafe(data)
}

func (ms ExportRequest) ReturnToPool() {
	ms.orig.ReturnToVTPool()
}

// UnmarshalProto unmarshalls ExportRequest from proto bytes.
func (ms ExportRequest) UnmarshalProto(data []byte) error {
	if err := ms.orig.UnmarshalVT(data); err != nil {
		return err
	}
	otlp.MigrateTraces(ms.orig.ResourceSpans)
	return nil
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
	td, err := jsonUnmarshaler.UnmarshalTraces(data)
	if err != nil {
		return err
	}
	*ms.orig = *internal.GetOrigTraces(internal.Traces(td))
	return nil
}

func (ms ExportRequest) Traces() ptrace.Traces {
	return ptrace.Traces(internal.NewTraces(ms.orig, ms.state))
}
