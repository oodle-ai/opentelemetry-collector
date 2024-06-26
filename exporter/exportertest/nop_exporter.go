// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exportertest // import "github.com/oodle-ai/opentelemetry-collector/exporter/exportertest"

import (
	"context"

	"github.com/google/uuid"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/component/componenttest"
	"github.com/oodle-ai/opentelemetry-collector/consumer/consumertest"
	"github.com/oodle-ai/opentelemetry-collector/exporter"
)

var nopType = component.MustNewType("nop")

// NewNopCreateSettings returns a new nop settings for Create*Exporter functions.
func NewNopCreateSettings() exporter.CreateSettings {
	return exporter.CreateSettings{
		ID:                component.NewIDWithName(nopType, uuid.NewString()),
		TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		BuildInfo:         component.NewDefaultBuildInfo(),
	}
}

// NewNopFactory returns an exporter.Factory that constructs nop exporters.
func NewNopFactory() exporter.Factory {
	return exporter.NewFactory(
		nopType,
		func() component.Config { return &nopConfig{} },
		exporter.WithTraces(createTracesExporter, component.StabilityLevelStable),
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelStable),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelStable),
	)
}

func createTracesExporter(context.Context, exporter.CreateSettings, component.Config) (exporter.Traces, error) {
	return nopInstance, nil
}

func createMetricsExporter(context.Context, exporter.CreateSettings, component.Config) (exporter.Metrics, error) {
	return nopInstance, nil
}

func createLogsExporter(context.Context, exporter.CreateSettings, component.Config) (exporter.Logs, error) {
	return nopInstance, nil
}

type nopConfig struct{}

var nopInstance = &nopExporter{
	Consumer: consumertest.NewNop(),
}

// nopExporter stores consumed traces and metrics for testing purposes.
type nopExporter struct {
	component.StartFunc
	component.ShutdownFunc
	consumertest.Consumer
}

// NewNopBuilder returns an exporter.Builder that constructs nop receivers.
func NewNopBuilder() *exporter.Builder {
	nopFactory := NewNopFactory()
	return exporter.NewBuilder(
		map[component.ID]component.Config{component.NewID(nopType): nopFactory.CreateDefaultConfig()},
		map[component.Type]exporter.Factory{nopType: nopFactory})
}
