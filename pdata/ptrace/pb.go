// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ptrace // import "github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"

import (
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlptrace "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/trace/v1"
)

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (e *ProtoMarshaler) MarshalTraces(td Traces) ([]byte, error) {
	pb := internal.TracesToProto(internal.Traces(td))
	return pb.Marshal()
}

func (e *ProtoMarshaler) TracesSize(td Traces) int {
	pb := internal.TracesToProto(internal.Traces(td))
	return pb.Size()
}

type ProtoUnmarshaler struct{}

func (d *ProtoUnmarshaler) UnmarshalTraces(buf []byte) (Traces, error) {
	pb := otlptrace.TracesData{}
	err := pb.Unmarshal(buf)
	return Traces(internal.TracesFromProto(pb)), err
}
