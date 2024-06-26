// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivertest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/consumer/consumertest"
)

var nopType = component.MustNewType("nop")

func TestNewNopFactory(t *testing.T) {
	factory := NewNopFactory()
	require.NotNil(t, factory)
	assert.Equal(t, nopType, factory.Type())
	cfg := factory.CreateDefaultConfig()
	assert.Equal(t, &nopConfig{}, cfg)

	traces, err := factory.CreateTracesReceiver(context.Background(), NewNopCreateSettings(), cfg, consumertest.NewNop())
	require.NoError(t, err)
	assert.NoError(t, traces.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, traces.Shutdown(context.Background()))

	metrics, err := factory.CreateMetricsReceiver(context.Background(), NewNopCreateSettings(), cfg, consumertest.NewNop())
	require.NoError(t, err)
	assert.NoError(t, metrics.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, metrics.Shutdown(context.Background()))

	logs, err := factory.CreateLogsReceiver(context.Background(), NewNopCreateSettings(), cfg, consumertest.NewNop())
	require.NoError(t, err)
	assert.NoError(t, logs.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, logs.Shutdown(context.Background()))
}

func TestNewNopBuilder(t *testing.T) {
	builder := NewNopBuilder()
	require.NotNil(t, builder)

	factory := NewNopFactory()
	cfg := factory.CreateDefaultConfig()
	set := NewNopCreateSettings()
	set.ID = component.NewID(nopType)

	traces, err := factory.CreateTracesReceiver(context.Background(), set, cfg, consumertest.NewNop())
	require.NoError(t, err)
	bTraces, err := builder.CreateTraces(context.Background(), set, consumertest.NewNop())
	require.NoError(t, err)
	assert.IsType(t, traces, bTraces)

	metrics, err := factory.CreateMetricsReceiver(context.Background(), set, cfg, consumertest.NewNop())
	require.NoError(t, err)
	bMetrics, err := builder.CreateMetrics(context.Background(), set, consumertest.NewNop())
	require.NoError(t, err)
	assert.IsType(t, metrics, bMetrics)

	logs, err := factory.CreateLogsReceiver(context.Background(), set, cfg, consumertest.NewNop())
	require.NoError(t, err)
	bLogs, err := builder.CreateLogs(context.Background(), set, consumertest.NewNop())
	require.NoError(t, err)
	assert.IsType(t, logs, bLogs)
}
