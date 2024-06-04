// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consumer // import "github.com/oodle-ai/opentelemetry-collector/consumer"

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
)

// Logs is an interface that receives plog.Logs, processes it
// as needed, and sends it to the next processing node if any or to the destination.
type Logs interface {
	baseConsumer
	// ConsumeLogs receives plog.Logs for consumption.
	ConsumeLogs(ctx context.Context, ld plog.Logs) error
}

// ConsumeLogsFunc is a helper function that is similar to ConsumeLogs.
type ConsumeLogsFunc func(ctx context.Context, ld plog.Logs) error

// ConsumeLogs calls f(ctx, ld).
func (f ConsumeLogsFunc) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	return f(ctx, ld)
}

type baseLogs struct {
	*baseImpl
	ConsumeLogsFunc
}

// NewLogs returns a Logs configured with the provided options.
func NewLogs(consume ConsumeLogsFunc, options ...Option) (Logs, error) {
	if consume == nil {
		return nil, errNilFunc
	}
	return &baseLogs{
		baseImpl:        newBaseImpl(options...),
		ConsumeLogsFunc: consume,
	}, nil
}
