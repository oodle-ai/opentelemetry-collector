// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package connectortest // import "github.com/oodle-ai/opentelemetry-collector/connector/connectortest"

import (
	"context"

	"github.com/google/uuid"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/connector"
	"github.com/oodle-ai/opentelemetry-collector/consumer"
	"github.com/oodle-ai/opentelemetry-collector/consumer/consumertest"
)

var nopType = component.MustNewType("nop")

// NewNopCreateSettings returns a new nop settings for Create* functions.
func NewNopCreateSettings() connector.CreateSettings {
	return connector.CreateSettings{
		ID:                component.NewIDWithName(nopType, uuid.NewString()),
		TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		BuildInfo:         component.NewDefaultBuildInfo(),
	}
}

type nopConfig struct{}

// NewNopFactory returns a connector.Factory that constructs nop processors.
func NewNopFactory() connector.Factory {
	return connector.NewFactory(
		nopType,
		func() component.Config {
			return &nopConfig{}
		},
		connector.WithTracesToTraces(createTracesToTracesConnector, component.StabilityLevelDevelopment),
		connector.WithTracesToMetrics(createTracesToMetricsConnector, component.StabilityLevelDevelopment),
		connector.WithTracesToLogs(createTracesToLogsConnector, component.StabilityLevelDevelopment),
		connector.WithMetricsToTraces(createMetricsToTracesConnector, component.StabilityLevelDevelopment),
		connector.WithMetricsToMetrics(createMetricsToMetricsConnector, component.StabilityLevelDevelopment),
		connector.WithMetricsToLogs(createMetricsToLogsConnector, component.StabilityLevelDevelopment),
		connector.WithLogsToTraces(createLogsToTracesConnector, component.StabilityLevelDevelopment),
		connector.WithLogsToMetrics(createLogsToMetricsConnector, component.StabilityLevelDevelopment),
		connector.WithLogsToLogs(createLogsToLogsConnector, component.StabilityLevelDevelopment),
	)
}

func createTracesToTracesConnector(context.Context, connector.CreateSettings, component.Config, consumer.Traces) (connector.Traces, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

func createTracesToMetricsConnector(context.Context, connector.CreateSettings, component.Config, consumer.Metrics) (connector.Traces, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

func createTracesToLogsConnector(context.Context, connector.CreateSettings, component.Config, consumer.Logs) (connector.Traces, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

func createMetricsToTracesConnector(context.Context, connector.CreateSettings, component.Config, consumer.Traces) (connector.Metrics, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

func createMetricsToMetricsConnector(context.Context, connector.CreateSettings, component.Config, consumer.Metrics) (connector.Metrics, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

func createMetricsToLogsConnector(context.Context, connector.CreateSettings, component.Config, consumer.Logs) (connector.Metrics, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

func createLogsToTracesConnector(context.Context, connector.CreateSettings, component.Config, consumer.Traces) (connector.Logs, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

func createLogsToMetricsConnector(context.Context, connector.CreateSettings, component.Config, consumer.Metrics) (connector.Logs, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

func createLogsToLogsConnector(context.Context, connector.CreateSettings, component.Config, consumer.Logs) (connector.Logs, error) {
	return &nopConnector{Consumer: consumertest.NewNop()}, nil
}

// nopConnector stores consumed traces and metrics for testing purposes.
type nopConnector struct {
	component.StartFunc
	component.ShutdownFunc
	consumertest.Consumer
}

// NewNopBuilder returns a connector.Builder that constructs nop receivers.
func NewNopBuilder() *connector.Builder {
	nopFactory := NewNopFactory()
	// Use a different ID than receivertest and exportertest to avoid ambiguous
	// configuration scenarios. Ambiguous IDs are detected in the 'otelcol' package,
	// but lower level packages such as 'service' assume that IDs are disambiguated.
	connID := component.NewIDWithName(nopType, "conn")
	return connector.NewBuilder(
		map[component.ID]component.Config{connID: nopFactory.CreateDefaultConfig()},
		map[component.Type]connector.Factory{nopType: nopFactory})
}
