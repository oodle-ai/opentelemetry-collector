// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package memorylimiterprocessor // import "github.com/oodle-ai/opentelemetry-collector/processor/memorylimiterprocessor"

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/internal/memorylimiter"
	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"
	"github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"
	"github.com/oodle-ai/opentelemetry-collector/processor"
	"github.com/oodle-ai/opentelemetry-collector/processor/processorhelper"
)

type memoryLimiterProcessor struct {
	memlimiter *memorylimiter.MemoryLimiter
	obsrep     *processorhelper.ObsReport
}

// newMemoryLimiter returns a new memorylimiter processor.
func newMemoryLimiterProcessor(set processor.CreateSettings, cfg *Config) (*memoryLimiterProcessor, error) {
	ml, err := memorylimiter.NewMemoryLimiter(cfg, set.Logger)
	if err != nil {
		return nil, err
	}
	obsrep, err := processorhelper.NewObsReport(processorhelper.ObsReportSettings{
		ProcessorID:             set.ID,
		ProcessorCreateSettings: set,
	})
	if err != nil {
		return nil, err
	}

	p := &memoryLimiterProcessor{
		memlimiter: ml,
		obsrep:     obsrep,
	}

	return p, nil
}

func (p *memoryLimiterProcessor) start(ctx context.Context, host component.Host) error {
	return p.memlimiter.Start(ctx, host)
}

func (p *memoryLimiterProcessor) shutdown(ctx context.Context) error {
	return p.memlimiter.Shutdown(ctx)
}

func (p *memoryLimiterProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	numSpans := td.SpanCount()
	if p.memlimiter.MustRefuse() {
		// TODO: actually to be 100% sure that this is "refused" and not "dropped"
		// 	it is necessary to check the pipeline to see if this is directly connected
		// 	to a receiver (ie.: a receiver is on the call stack). For now it
		// 	assumes that the pipeline is properly configured and a receiver is on the
		// 	callstack and that the receiver will correctly retry the refused data again.
		p.obsrep.TracesRefused(ctx, numSpans)
		return td, memorylimiter.ErrDataRefused
	}

	// Even if the next consumer returns error record the data as accepted by
	// this processor.
	p.obsrep.TracesAccepted(ctx, numSpans)
	return td, nil
}

func (p *memoryLimiterProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	numDataPoints := md.DataPointCount()
	if p.memlimiter.MustRefuse() {
		// TODO: actually to be 100% sure that this is "refused" and not "dropped"
		// 	it is necessary to check the pipeline to see if this is directly connected
		// 	to a receiver (ie.: a receiver is on the call stack). For now it
		// 	assumes that the pipeline is properly configured and a receiver is on the
		// 	callstack.
		p.obsrep.MetricsRefused(ctx, numDataPoints)
		return md, memorylimiter.ErrDataRefused
	}

	// Even if the next consumer returns error record the data as accepted by
	// this processor.
	p.obsrep.MetricsAccepted(ctx, numDataPoints)
	return md, nil
}

func (p *memoryLimiterProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	numRecords := ld.LogRecordCount()
	if p.memlimiter.MustRefuse() {
		// TODO: actually to be 100% sure that this is "refused" and not "dropped"
		// 	it is necessary to check the pipeline to see if this is directly connected
		// 	to a receiver (ie.: a receiver is on the call stack). For now it
		// 	assumes that the pipeline is properly configured and a receiver is on the
		// 	callstack.
		p.obsrep.LogsRefused(ctx, numRecords)
		return ld, memorylimiter.ErrDataRefused
	}

	// Even if the next consumer returns error record the data as accepted by
	// this processor.
	p.obsrep.LogsAccepted(ctx, numRecords)
	return ld, nil
}
