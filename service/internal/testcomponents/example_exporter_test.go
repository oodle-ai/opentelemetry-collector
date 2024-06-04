// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testcomponents

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"
	"github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"
)

func TestExampleExporter(t *testing.T) {
	exp := &ExampleExporter{}
	host := componenttest.NewNopHost()
	assert.False(t, exp.Started())
	assert.NoError(t, exp.Start(context.Background(), host))
	assert.True(t, exp.Started())

	assert.Equal(t, 0, len(exp.Traces))
	assert.NoError(t, exp.ConsumeTraces(context.Background(), ptrace.Traces{}))
	assert.Equal(t, 1, len(exp.Traces))

	assert.Equal(t, 0, len(exp.Metrics))
	assert.NoError(t, exp.ConsumeMetrics(context.Background(), pmetric.Metrics{}))
	assert.Equal(t, 1, len(exp.Metrics))

	assert.Equal(t, 0, len(exp.Logs))
	assert.NoError(t, exp.ConsumeLogs(context.Background(), plog.Logs{}))
	assert.Equal(t, 1, len(exp.Logs))

	assert.False(t, exp.Stopped())
	assert.NoError(t, exp.Shutdown(context.Background()))
	assert.True(t, exp.Stopped())
}
