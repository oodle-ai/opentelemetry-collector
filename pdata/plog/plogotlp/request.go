// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package plogotlp // import "github.com/oodle-ai/opentelemetry-collector/pdata/plog/plogotlp"

import (
	"bytes"

	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlpcollectorlog "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/collector/logs/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal/json"
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal/otlp"
	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
)

var jsonUnmarshaler = &plog.JSONUnmarshaler{}

// ExportRequest represents the request for gRPC/HTTP client/server.
// It's a wrapper for plog.Logs data.
type ExportRequest struct {
	orig  *otlpcollectorlog.ExportLogsServiceRequest
	state *internal.State
}

// NewExportRequest returns an empty ExportRequest.
func NewExportRequest() ExportRequest {
	state := internal.StateMutable
	return ExportRequest{
		orig:  otlpcollectorlog.ExportLogsServiceRequestFromVTPool(),
		state: &state,
	}
}

// NewExportRequestFromLogs returns a ExportRequest from plog.Logs.
// Because ExportRequest is a wrapper for plog.Logs,
// any changes to the provided Logs struct will be reflected in the ExportRequest and vice versa.
func NewExportRequestFromLogs(ld plog.Logs) ExportRequest {
	return ExportRequest{
		orig:  internal.GetOrigLogs(internal.Logs(ld)),
		state: internal.GetLogsState(internal.Logs(ld)),
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
	otlp.MigrateLogs(ms.orig.ResourceLogs)
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
	ld, err := jsonUnmarshaler.UnmarshalLogs(data)
	if err != nil {
		return err
	}
	*ms.orig = *internal.GetOrigLogs(internal.Logs(ld))
	return nil
}

func (ms ExportRequest) Logs() plog.Logs {
	return plog.Logs(internal.NewLogs(ms.orig, ms.state))
}
