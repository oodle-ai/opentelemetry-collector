// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processorhelper

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/processor"
)

var (
	processorID = component.MustNewID("fakeProcessor")
)

func TestProcessorTraceData(t *testing.T) {
	testTelemetry(t, processorID, func(t *testing.T, tt componenttest.TestTelemetry) {
		const acceptedSpans = 27
		const refusedSpans = 19
		const droppedSpans = 13
		obsrep, err := newObsReport(ObsReportSettings{
			ProcessorID:             processorID,
			ProcessorCreateSettings: processor.CreateSettings{ID: processorID, TelemetrySettings: tt.TelemetrySettings(), BuildInfo: component.NewDefaultBuildInfo()},
		})
		require.NoError(t, err)
		obsrep.TracesAccepted(context.Background(), acceptedSpans)
		obsrep.TracesRefused(context.Background(), refusedSpans)
		obsrep.TracesDropped(context.Background(), droppedSpans)

		require.NoError(t, tt.CheckProcessorTraces(acceptedSpans, refusedSpans, droppedSpans))
	})
}

func TestProcessorMetricsData(t *testing.T) {
	testTelemetry(t, processorID, func(t *testing.T, tt componenttest.TestTelemetry) {
		const acceptedPoints = 29
		const refusedPoints = 11
		const droppedPoints = 17

		obsrep, err := newObsReport(ObsReportSettings{
			ProcessorID:             processorID,
			ProcessorCreateSettings: processor.CreateSettings{ID: processorID, TelemetrySettings: tt.TelemetrySettings(), BuildInfo: component.NewDefaultBuildInfo()},
		})
		require.NoError(t, err)
		obsrep.MetricsAccepted(context.Background(), acceptedPoints)
		obsrep.MetricsRefused(context.Background(), refusedPoints)
		obsrep.MetricsDropped(context.Background(), droppedPoints)

		require.NoError(t, tt.CheckProcessorMetrics(acceptedPoints, refusedPoints, droppedPoints))
	})
}

func TestBuildProcessorCustomMetricName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "firstMeasure",
			want: "processor_test_type_firstMeasure",
		},
		{
			name: "secondMeasure",
			want: "processor_test_type_secondMeasure",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildCustomMetricName("test_type", tt.name)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProcessorLogRecords(t *testing.T) {
	testTelemetry(t, processorID, func(t *testing.T, tt componenttest.TestTelemetry) {
		const acceptedRecords = 29
		const refusedRecords = 11
		const droppedRecords = 17

		obsrep, err := newObsReport(ObsReportSettings{
			ProcessorID:             processorID,
			ProcessorCreateSettings: processor.CreateSettings{ID: processorID, TelemetrySettings: tt.TelemetrySettings(), BuildInfo: component.NewDefaultBuildInfo()},
		})
		require.NoError(t, err)
		obsrep.LogsAccepted(context.Background(), acceptedRecords)
		obsrep.LogsRefused(context.Background(), refusedRecords)
		obsrep.LogsDropped(context.Background(), droppedRecords)

		require.NoError(t, tt.CheckProcessorLogs(acceptedRecords, refusedRecords, droppedRecords))
	})
}

func TestCheckProcessorTracesViews(t *testing.T) {
	tt, err := componenttest.SetupTelemetry(processorID)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, tt.Shutdown(context.Background())) })

	por, err := NewObsReport(ObsReportSettings{
		ProcessorID:             processorID,
		ProcessorCreateSettings: processor.CreateSettings{ID: processorID, TelemetrySettings: tt.TelemetrySettings(), BuildInfo: component.NewDefaultBuildInfo()},
	})
	assert.NoError(t, err)

	por.TracesAccepted(context.Background(), 7)
	por.TracesRefused(context.Background(), 8)
	por.TracesDropped(context.Background(), 9)

	assert.NoError(t, tt.CheckProcessorTraces(7, 8, 9))
	assert.Error(t, tt.CheckProcessorTraces(0, 0, 0))
	assert.Error(t, tt.CheckProcessorTraces(7, 0, 0))
	assert.Error(t, tt.CheckProcessorTraces(7, 8, 0))
	assert.Error(t, tt.CheckProcessorTraces(7, 0, 9))
	assert.Error(t, tt.CheckProcessorTraces(0, 8, 0))
	assert.Error(t, tt.CheckProcessorTraces(0, 8, 9))
	assert.Error(t, tt.CheckProcessorTraces(0, 0, 9))
}

func TestCheckProcessorMetricsViews(t *testing.T) {
	tt, err := componenttest.SetupTelemetry(processorID)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, tt.Shutdown(context.Background())) })

	por, err := NewObsReport(ObsReportSettings{
		ProcessorID:             processorID,
		ProcessorCreateSettings: processor.CreateSettings{ID: processorID, TelemetrySettings: tt.TelemetrySettings(), BuildInfo: component.NewDefaultBuildInfo()},
	})
	assert.NoError(t, err)

	por.MetricsAccepted(context.Background(), 7)
	por.MetricsRefused(context.Background(), 8)
	por.MetricsDropped(context.Background(), 9)

	assert.NoError(t, tt.CheckProcessorMetrics(7, 8, 9))
	assert.Error(t, tt.CheckProcessorMetrics(0, 0, 0))
	assert.Error(t, tt.CheckProcessorMetrics(7, 0, 0))
	assert.Error(t, tt.CheckProcessorMetrics(7, 8, 0))
	assert.Error(t, tt.CheckProcessorMetrics(7, 0, 9))
	assert.Error(t, tt.CheckProcessorMetrics(0, 8, 0))
	assert.Error(t, tt.CheckProcessorMetrics(0, 8, 9))
	assert.Error(t, tt.CheckProcessorMetrics(0, 0, 9))
}

func TestCheckProcessorLogViews(t *testing.T) {
	tt, err := componenttest.SetupTelemetry(processorID)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, tt.Shutdown(context.Background())) })

	por, err := NewObsReport(ObsReportSettings{
		ProcessorID:             processorID,
		ProcessorCreateSettings: processor.CreateSettings{ID: processorID, TelemetrySettings: tt.TelemetrySettings(), BuildInfo: component.NewDefaultBuildInfo()},
	})
	assert.NoError(t, err)

	por.LogsAccepted(context.Background(), 7)
	por.LogsRefused(context.Background(), 8)
	por.LogsDropped(context.Background(), 9)

	assert.NoError(t, tt.CheckProcessorLogs(7, 8, 9))
	assert.Error(t, tt.CheckProcessorLogs(0, 0, 0))
	assert.Error(t, tt.CheckProcessorLogs(7, 0, 0))
	assert.Error(t, tt.CheckProcessorLogs(7, 8, 0))
	assert.Error(t, tt.CheckProcessorLogs(7, 0, 9))
	assert.Error(t, tt.CheckProcessorLogs(0, 8, 0))
	assert.Error(t, tt.CheckProcessorLogs(0, 8, 9))
	assert.Error(t, tt.CheckProcessorLogs(0, 0, 9))
}

func testTelemetry(t *testing.T, id component.ID, testFunc func(t *testing.T, tt componenttest.TestTelemetry)) {
	tt, err := componenttest.SetupTelemetry(id)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, tt.Shutdown(context.Background())) })

	testFunc(t, tt)
}
