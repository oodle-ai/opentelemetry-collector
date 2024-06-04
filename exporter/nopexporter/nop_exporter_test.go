// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nopexporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/exporter/exportertest"
	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"
	"github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"
)

func TestNewNopFactory(t *testing.T) {
	factory := NewFactory()
	require.NotNil(t, factory)
	assert.Equal(t, component.MustNewType("nop"), factory.Type())
	cfg := factory.CreateDefaultConfig()
	assert.Equal(t, &struct{}{}, cfg)

	traces, err := factory.CreateTracesExporter(context.Background(), exportertest.NewNopCreateSettings(), cfg)
	require.NoError(t, err)
	assert.NoError(t, traces.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, traces.ConsumeTraces(context.Background(), ptrace.NewTraces()))
	assert.NoError(t, traces.Shutdown(context.Background()))

	metrics, err := factory.CreateMetricsExporter(context.Background(), exportertest.NewNopCreateSettings(), cfg)
	require.NoError(t, err)
	assert.NoError(t, metrics.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, metrics.ConsumeMetrics(context.Background(), pmetric.NewMetrics()))
	assert.NoError(t, metrics.Shutdown(context.Background()))

	logs, err := factory.CreateLogsExporter(context.Background(), exportertest.NewNopCreateSettings(), cfg)
	require.NoError(t, err)
	assert.NoError(t, logs.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, logs.ConsumeLogs(context.Background(), plog.NewLogs()))
	assert.NoError(t, logs.Shutdown(context.Background()))
}
