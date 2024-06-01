// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testcomponents // import "github.com/oodle-ai/opentelemetry-collector/service/internal/testcomponents"

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/consumer"
	"github.com/oodle-ai/opentelemetry-collector/exporter"
	"github.com/oodle-ai/opentelemetry-collector/pdata/plog"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pmetric"
	"github.com/oodle-ai/opentelemetry-collector/pdata/ptrace"
)

var testType = component.MustNewType("exampleexporter")

const stability = component.StabilityLevelDevelopment

// ExampleExporterFactory is factory for ExampleExporter.
var ExampleExporterFactory = exporter.NewFactory(
	testType,
	createExporterDefaultConfig,
	exporter.WithTraces(createTracesExporter, stability),
	exporter.WithMetrics(createMetricsExporter, stability),
	exporter.WithLogs(createLogsExporter, stability),
)

func createExporterDefaultConfig() component.Config {
	return &struct{}{}
}

func createTracesExporter(context.Context, exporter.CreateSettings, component.Config) (exporter.Traces, error) {
	return &ExampleExporter{}, nil
}

func createMetricsExporter(context.Context, exporter.CreateSettings, component.Config) (exporter.Metrics, error) {
	return &ExampleExporter{}, nil
}

func createLogsExporter(context.Context, exporter.CreateSettings, component.Config) (exporter.Logs, error) {
	return &ExampleExporter{}, nil
}

// ExampleExporter stores consumed traces and metrics for testing purposes.
type ExampleExporter struct {
	componentState
	Traces  []ptrace.Traces
	Metrics []pmetric.Metrics
	Logs    []plog.Logs
}

// ConsumeTraces receives ptrace.Traces for processing by the consumer.Traces.
func (exp *ExampleExporter) ConsumeTraces(_ context.Context, td ptrace.Traces) error {
	exp.Traces = append(exp.Traces, td)
	return nil
}

// ConsumeMetrics receives pmetric.Metrics for processing by the Metrics.
func (exp *ExampleExporter) ConsumeMetrics(_ context.Context, md pmetric.Metrics) error {
	exp.Metrics = append(exp.Metrics, md)
	return nil
}

// ConsumeLogs receives plog.Logs for processing by the Logs.
func (exp *ExampleExporter) ConsumeLogs(_ context.Context, ld plog.Logs) error {
	exp.Logs = append(exp.Logs, ld)
	return nil
}

func (exp *ExampleExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
