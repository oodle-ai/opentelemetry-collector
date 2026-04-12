// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package plog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	goproto "google.golang.org/protobuf/proto"

	otlpcollectorlog "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/collector/logs/v1"
	otlplogs "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/logs/v1"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

func TestLogRecordCount(t *testing.T) {
	logs := NewLogs()
	assert.EqualValues(t, 0, logs.LogRecordCount())

	rl := logs.ResourceLogs().AppendEmpty()
	assert.EqualValues(t, 0, logs.LogRecordCount())

	ill := rl.ScopeLogs().AppendEmpty()
	assert.EqualValues(t, 0, logs.LogRecordCount())

	ill.LogRecords().AppendEmpty()
	assert.EqualValues(t, 1, logs.LogRecordCount())

	rms := logs.ResourceLogs()
	rms.EnsureCapacity(3)
	rms.AppendEmpty().ScopeLogs().AppendEmpty()
	illl := rms.AppendEmpty().ScopeLogs().AppendEmpty().LogRecords()
	for i := 0; i < 5; i++ {
		illl.AppendEmpty()
	}
	// 5 + 1 (from rms.At(0) initialized first)
	assert.EqualValues(t, 6, logs.LogRecordCount())
}

func TestLogRecordCountWithEmpty(t *testing.T) {
	assert.Zero(t, NewLogs().LogRecordCount())
	assert.Zero(t, newLogs(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{{}},
	}).LogRecordCount())
	assert.Zero(t, newLogs(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{
			{
				ScopeLogs: []*otlplogs.ScopeLogs{{}},
			},
		},
	}).LogRecordCount())
	assert.Equal(t, 1, newLogs(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{
			{
				ScopeLogs: []*otlplogs.ScopeLogs{
					{
						LogRecords: []*otlplogs.LogRecord{{}},
					},
				},
			},
		},
	}).LogRecordCount())
}

func TestToFromLogOtlp(t *testing.T) {
	otlp := &otlpcollectorlog.ExportLogsServiceRequest{}
	logs := newLogs(otlp)
	assert.EqualValues(t, NewLogs(), logs)
	assert.EqualValues(t, otlp, logs.getOrig())
}

func TestResourceLogsWireCompatibility(t *testing.T) {
	// Verifies a full protobuf marshal → unmarshal → marshal → unmarshal round trip
	// for ExportLogsServiceRequest (vtproto / golang protobuf generated types).

	logs := NewLogs()
	fillTestResourceLogsSlice(logs.ResourceLogs())

	wire1, err := goproto.Marshal(logs.getOrig())
	assert.NoError(t, err)
	assert.NotEmpty(t, wire1)

	var decoded otlpcollectorlog.ExportLogsServiceRequest
	err = goproto.Unmarshal(wire1, &decoded)
	assert.NoError(t, err)

	wire2, err := goproto.Marshal(&decoded)
	assert.NoError(t, err)
	assert.NotEmpty(t, wire2)

	var roundTrip otlpcollectorlog.ExportLogsServiceRequest
	err = goproto.Unmarshal(wire2, &roundTrip)
	assert.NoError(t, err)

	assert.True(t, goproto.Equal(logs.getOrig(), &roundTrip))
}

func TestLogsCopyTo(t *testing.T) {
	logs := NewLogs()
	fillTestResourceLogsSlice(logs.ResourceLogs())
	logsCopy := NewLogs()
	logs.CopyTo(logsCopy)
	assert.EqualValues(t, logs, logsCopy)
}

func TestReadOnlyLogsInvalidUsage(t *testing.T) {
	logs := NewLogs()
	assert.False(t, logs.IsReadOnly())
	res := logs.ResourceLogs().AppendEmpty().Resource()
	res.Attributes().PutStr("k1", "v1")
	logs.MarkReadOnly()
	assert.True(t, logs.IsReadOnly())
	assert.Panics(t, func() { res.Attributes().PutStr("k2", "v2") })
}

func BenchmarkLogsUsage(b *testing.B) {
	logs := NewLogs()
	fillTestResourceLogsSlice(logs.ResourceLogs())

	ts := pcommon.NewTimestampFromTime(time.Now())

	b.ReportAllocs()
	b.ResetTimer()

	for bb := 0; bb < b.N; bb++ {
		for i := 0; i < logs.ResourceLogs().Len(); i++ {
			rl := logs.ResourceLogs().At(i)
			res := rl.Resource()
			res.Attributes().PutStr("foo", "bar")
			v, ok := res.Attributes().Get("foo")
			assert.True(b, ok)
			assert.Equal(b, "bar", v.Str())
			v.SetStr("new-bar")
			assert.Equal(b, "new-bar", v.Str())
			res.Attributes().Remove("foo")
			for j := 0; j < rl.ScopeLogs().Len(); j++ {
				sl := rl.ScopeLogs().At(j)
				sl.Scope().SetName("new_test_name")
				assert.Equal(b, "new_test_name", sl.Scope().Name())
				for k := 0; k < sl.LogRecords().Len(); k++ {
					lr := sl.LogRecords().At(k)
					lr.Body().SetStr("new_body")
					assert.Equal(b, "new_body", lr.Body().Str())
					lr.SetTimestamp(ts)
					assert.Equal(b, ts, lr.Timestamp())
				}
				lr := sl.LogRecords().AppendEmpty()
				lr.Body().SetStr("another_log_record")
				lr.SetTimestamp(ts)
				lr.SetObservedTimestamp(ts)
				lr.SetSeverityText("info")
				lr.SetSeverityNumber(SeverityNumberInfo)
				lr.Attributes().PutStr("foo", "bar")
				lr.SetSpanID([8]byte{1, 2, 3, 4, 5, 6, 7, 8})
				sl.LogRecords().RemoveIf(func(lr LogRecord) bool {
					return lr.Body().Str() == "another_log_record"
				})
			}
		}
	}
}
