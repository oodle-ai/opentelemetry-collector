// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processorhelper // import "github.com/oodle-ai/opentelemetry-collector/processor/processorhelper"

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/trace"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/consumer"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"
	"github.com/oodle-ai/opentelemetry-collector/processor"
)

// ProcessMetricsFunc is a helper function that processes the incoming data and returns the data to be sent to the next component.
// If error is returned then returned data are ignored. It MUST not call the next component.
type ProcessMetricsFunc func(context.Context, pmetric.Metrics) (pmetric.Metrics, error)

type metricsProcessor struct {
	component.StartFunc
	component.ShutdownFunc
	consumer.Metrics
}

// NewMetricsProcessor creates a processor.Metrics that ensure context propagation and the right tags are set.
func NewMetricsProcessor(
	_ context.Context,
	set processor.CreateSettings,
	_ component.Config,
	nextConsumer consumer.Metrics,
	metricsFunc ProcessMetricsFunc,
	options ...Option,
) (processor.Metrics, error) {
	// TODO: Add observability metrics support
	if metricsFunc == nil {
		return nil, errors.New("nil metricsFunc")
	}

	eventOptions := spanAttributes(set.ID)
	bs := fromOptions(options)
	metricsConsumer, err := consumer.NewMetrics(func(ctx context.Context, md pmetric.Metrics) error {
		span := trace.SpanFromContext(ctx)
		span.AddEvent("Start processing.", eventOptions)
		var err error
		md, err = metricsFunc(ctx, md)
		span.AddEvent("End processing.", eventOptions)
		if err != nil {
			if errors.Is(err, ErrSkipProcessingData) {
				return nil
			}
			return err
		}
		return nextConsumer.ConsumeMetrics(ctx, md)
	}, bs.consumerOptions...)
	if err != nil {
		return nil, err
	}

	return &metricsProcessor{
		StartFunc:    bs.StartFunc,
		ShutdownFunc: bs.ShutdownFunc,
		Metrics:      metricsConsumer,
	}, nil
}
