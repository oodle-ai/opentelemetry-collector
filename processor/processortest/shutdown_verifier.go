// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processortest // import "github.com/oodle-ai/opentelemetry-collector/processor/processortest"

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/consumer/consumertest"
	"github.com/oodle-ai/opentelemetry-collector/pdata/testdata"
	"github.com/oodle-ai/opentelemetry-collector/processor"
)

func verifyTracesDoesNotProduceAfterShutdown(t *testing.T, factory processor.Factory, cfg component.Config) {
	// Create a proc and output its produce to a sink.
	nextSink := new(consumertest.TracesSink)
	proc, err := factory.CreateTracesProcessor(context.Background(), NewNopCreateSettings(), cfg, nextSink)
	if errors.Is(err, component.ErrDataTypeIsNotSupported) {
		return
	}
	require.NoError(t, err)
	assert.NoError(t, proc.Start(context.Background(), componenttest.NewNopHost()))

	// Send some traces to the proc.
	const generatedCount = 10
	for i := 0; i < generatedCount; i++ {
		require.NoError(t, proc.ConsumeTraces(context.Background(), testdata.GenerateTraces(1)))
	}

	// Now shutdown the proc.
	assert.NoError(t, proc.Shutdown(context.Background()))

	// The Shutdown() is done. It means the proc must have sent everything we
	// gave it to the next sink.
	assert.EqualValues(t, generatedCount, nextSink.SpanCount())
}

func verifyLogsDoesNotProduceAfterShutdown(t *testing.T, factory processor.Factory, cfg component.Config) {
	// Create a proc and output its produce to a sink.
	nextSink := new(consumertest.LogsSink)
	proc, err := factory.CreateLogsProcessor(context.Background(), NewNopCreateSettings(), cfg, nextSink)
	if errors.Is(err, component.ErrDataTypeIsNotSupported) {
		return
	}
	require.NoError(t, err)
	assert.NoError(t, proc.Start(context.Background(), componenttest.NewNopHost()))

	// Send some logs to the proc.
	const generatedCount = 10
	for i := 0; i < generatedCount; i++ {
		require.NoError(t, proc.ConsumeLogs(context.Background(), testdata.GenerateLogs(1)))
	}

	// Now shutdown the proc.
	assert.NoError(t, proc.Shutdown(context.Background()))

	// The Shutdown() is done. It means the proc must have sent everything we
	// gave it to the next sink.
	assert.EqualValues(t, generatedCount, nextSink.LogRecordCount())
}

func verifyMetricsDoesNotProduceAfterShutdown(t *testing.T, factory processor.Factory, cfg component.Config) {
	// Create a proc and output its produce to a sink.
	nextSink := new(consumertest.MetricsSink)
	proc, err := factory.CreateMetricsProcessor(context.Background(), NewNopCreateSettings(), cfg, nextSink)
	if errors.Is(err, component.ErrDataTypeIsNotSupported) {
		return
	}
	require.NoError(t, err)
	assert.NoError(t, proc.Start(context.Background(), componenttest.NewNopHost()))

	// Send some metrics to the proc. testdata.GenerateMetrics creates metrics with 2 data points each.
	const generatedCount = 10
	for i := 0; i < generatedCount; i++ {
		require.NoError(t, proc.ConsumeMetrics(context.Background(), testdata.GenerateMetrics(1)))
	}

	// Now shutdown the proc.
	assert.NoError(t, proc.Shutdown(context.Background()))

	// The Shutdown() is done. It means the proc must have sent everything we
	// gave it to the next sink.
	assert.EqualValues(t, generatedCount*2, nextSink.DataPointCount())
}

// VerifyShutdown verifies the processor doesn't produce telemetry data after shutdown.
func VerifyShutdown(t *testing.T, factory processor.Factory, cfg component.Config) {
	verifyTracesDoesNotProduceAfterShutdown(t, factory, cfg)
	verifyLogsDoesNotProduceAfterShutdown(t, factory, cfg)
	verifyMetricsDoesNotProduceAfterShutdown(t, factory, cfg)
}
