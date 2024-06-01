// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processortest // import "github.com/oodle-ai/opentelemetry-collector/processor/processortest"

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/consumer"
	"github.com/oodle-ai/opentelemetry-collector/consumer/consumertest"
	"github.com/oodle-ai/opentelemetry-collector/processor"
)

// NewUnhealthyProcessorCreateSettings returns a new nop settings for Create*Processor functions.
func NewUnhealthyProcessorCreateSettings() processor.CreateSettings {
	return processor.CreateSettings{
		TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		BuildInfo:         component.NewDefaultBuildInfo(),
	}
}

// NewUnhealthyProcessorFactory returns a component.ProcessorFactory that constructs nop processors.
func NewUnhealthyProcessorFactory() processor.Factory {
	return processor.NewFactory(
		component.MustNewType("unhealthy"),
		func() component.Config {
			return &struct{}{}
		},
		processor.WithTraces(createUnhealthyTracesProcessor, component.StabilityLevelStable),
		processor.WithMetrics(createUnhealthyMetricsProcessor, component.StabilityLevelStable),
		processor.WithLogs(createUnhealthyLogsProcessor, component.StabilityLevelStable),
	)
}

func createUnhealthyTracesProcessor(_ context.Context, set processor.CreateSettings, _ component.Config, _ consumer.Traces) (processor.Traces, error) {
	return &unhealthyProcessor{
		Consumer:  consumertest.NewNop(),
		telemetry: set.TelemetrySettings,
	}, nil
}

func createUnhealthyMetricsProcessor(_ context.Context, set processor.CreateSettings, _ component.Config, _ consumer.Metrics) (processor.Metrics, error) {
	return &unhealthyProcessor{
		Consumer:  consumertest.NewNop(),
		telemetry: set.TelemetrySettings,
	}, nil
}

func createUnhealthyLogsProcessor(_ context.Context, set processor.CreateSettings, _ component.Config, _ consumer.Logs) (processor.Logs, error) {
	return &unhealthyProcessor{
		Consumer:  consumertest.NewNop(),
		telemetry: set.TelemetrySettings,
	}, nil
}

type unhealthyProcessor struct {
	component.StartFunc
	component.ShutdownFunc
	consumertest.Consumer
	telemetry component.TelemetrySettings
}

func (p unhealthyProcessor) Start(context.Context, component.Host) error {
	go func() {
		p.telemetry.ReportStatus(component.NewStatusEvent(component.StatusRecoverableError))
	}()
	return nil
}
