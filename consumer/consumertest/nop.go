// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumertest // import "github.com/oodle-ai/opentelemetry-collector/consumer/consumertest"

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"
	"github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"
)

// NewNop returns a Consumer that just drops all received data and returns no error.
func NewNop() Consumer {
	return &baseConsumer{
		ConsumeTracesFunc:  func(context.Context, ptrace.Traces) error { return nil },
		ConsumeMetricsFunc: func(context.Context, pmetric.Metrics) error { return nil },
		ConsumeLogsFunc:    func(context.Context, plog.Logs) error { return nil },
	}
}
