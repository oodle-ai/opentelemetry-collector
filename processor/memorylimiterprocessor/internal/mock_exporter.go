// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/oodle-ai/opentelemetry-collector/processor/memorylimiterprocessor/internal"

import (
	"context"
	"sync/atomic"

	"github.com/oodle-ai/opentelemetry-collector/consumer"
	"github.com/oodle-ai/opentelemetry-collector/consumer/consumertest"
	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
)

type MockExporter struct {
	destAvailable     atomic.Bool
	acceptedLogCount  atomic.Int64
	deliveredLogCount atomic.Int64
	Logs              consumertest.LogsSink
}

var _ consumer.Logs = (*MockExporter)(nil)

func (e *MockExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{}
}

func (e *MockExporter) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	e.acceptedLogCount.Add(int64(ld.LogRecordCount()))

	if e.destAvailable.Load() {
		// Destination is available, immediately deliver.
		e.deliveredLogCount.Add(int64(ld.LogRecordCount()))
	} else {
		// Destination is not available. Queue the logs in the exporter.
		return e.Logs.ConsumeLogs(ctx, ld)
	}
	return nil
}

func (e *MockExporter) SetDestAvailable(available bool) {
	if available {
		// Pretend we delivered all queued accepted logs.
		e.deliveredLogCount.Add(int64(e.Logs.LogRecordCount()))

		// Get rid of the delivered logs so that memory can be collected.
		e.Logs.Reset()

		// Now mark destination available so that subsequent ConsumeLogs
		// don't queue the logs anymore.
		e.destAvailable.Store(true)
	} else {
		e.destAvailable.Store(false)
	}
}

func (e *MockExporter) AcceptedLogCount() int {
	return int(e.acceptedLogCount.Load())
}

func (e *MockExporter) DeliveredLogCount() int {
	return int(e.deliveredLogCount.Load())
}

func NewMockExporter() *MockExporter {
	return &MockExporter{
		destAvailable:     atomic.Bool{},
		acceptedLogCount:  atomic.Int64{},
		deliveredLogCount: atomic.Int64{},
		Logs:              consumertest.LogsSink{},
	}
}
