// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package plog // import "github.com/oodle-ai/opentelemetry-collector/pdata/plog"

import (
	"github.com/oodle-ai/opentelemetry-collector/pdata/internal"
	otlplogs "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/logs/v1"
)

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (e *ProtoMarshaler) MarshalLogs(ld Logs) ([]byte, error) {
	pb := internal.LogsToProto(internal.Logs(ld))
	return pb.MarshalVT()
}

func (e *ProtoMarshaler) LogsSize(ld Logs) int {
	pb := internal.LogsToProto(internal.Logs(ld))
	return pb.SizeVT()
}

var _ Unmarshaler = (*ProtoUnmarshaler)(nil)

type ProtoUnmarshaler struct{}

func (d *ProtoUnmarshaler) UnmarshalLogs(buf []byte) (Logs, error) {
	pb := otlplogs.LogsData{}
	err := pb.UnmarshalVT(buf)
	return Logs(internal.LogsFromProto(pb)), err
}
