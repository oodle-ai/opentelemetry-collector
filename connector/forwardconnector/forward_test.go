// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package forwardconnector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/connector/connectortest"
	"github.com/oodle-ai/opentelemetry-collector/consumer/consumertest"
	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"
	"github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"
)

func TestForward(t *testing.T) {
	f := NewFactory()
	cfg := f.CreateDefaultConfig()
	assert.Equal(t, &Config{}, cfg)

	ctx := context.Background()
	set := connectortest.NewNopCreateSettings()
	host := componenttest.NewNopHost()

	tracesSink := new(consumertest.TracesSink)
	tracesToTraces, err := f.CreateTracesToTraces(ctx, set, cfg, tracesSink)
	assert.NoError(t, err)
	assert.NotNil(t, tracesToTraces)

	metricsSink := new(consumertest.MetricsSink)
	metricsToMetrics, err := f.CreateMetricsToMetrics(ctx, set, cfg, metricsSink)
	assert.NoError(t, err)
	assert.NotNil(t, metricsToMetrics)

	logsSink := new(consumertest.LogsSink)
	logsToLogs, err := f.CreateLogsToLogs(ctx, set, cfg, logsSink)
	assert.NoError(t, err)
	assert.NotNil(t, logsToLogs)

	assert.NoError(t, tracesToTraces.Start(ctx, host))
	assert.NoError(t, metricsToMetrics.Start(ctx, host))
	assert.NoError(t, logsToLogs.Start(ctx, host))

	assert.NoError(t, tracesToTraces.ConsumeTraces(ctx, ptrace.NewTraces()))

	assert.NoError(t, metricsToMetrics.ConsumeMetrics(ctx, pmetric.NewMetrics()))
	assert.NoError(t, metricsToMetrics.ConsumeMetrics(ctx, pmetric.NewMetrics()))

	assert.NoError(t, logsToLogs.ConsumeLogs(ctx, plog.NewLogs()))
	assert.NoError(t, logsToLogs.ConsumeLogs(ctx, plog.NewLogs()))
	assert.NoError(t, logsToLogs.ConsumeLogs(ctx, plog.NewLogs()))

	assert.NoError(t, tracesToTraces.Shutdown(ctx))
	assert.NoError(t, metricsToMetrics.Shutdown(ctx))
	assert.NoError(t, logsToLogs.Shutdown(ctx))

	assert.Equal(t, 1, len(tracesSink.AllTraces()))
	assert.Equal(t, 2, len(metricsSink.AllMetrics()))
	assert.Equal(t, 3, len(logsSink.AllLogs()))
}
