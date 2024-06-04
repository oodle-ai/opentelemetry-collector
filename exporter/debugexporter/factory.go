// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package debugexporter // import "github.com/oodle-ai/opentelemetry-collector/exporter/debugexporter"

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/config/configtelemetry"
	"github.com/oodle-ai/opentelemetry-collector/consumer"
	"github.com/oodle-ai/opentelemetry-collector/exporter"
	"github.com/oodle-ai/opentelemetry-collector/exporter/debugexporter/internal/metadata"
	"github.com/oodle-ai/opentelemetry-collector/exporter/exporterhelper"
	"github.com/oodle-ai/opentelemetry-collector/exporter/internal/otlptext"
)

// The value of "type" key in configuration.
var componentType = component.MustNewType("debug")

const (
	defaultSamplingInitial    = 2
	defaultSamplingThereafter = 500
)

// NewFactory creates a factory for Debug exporter
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		componentType,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Verbosity:          configtelemetry.LevelBasic,
		SamplingInitial:    defaultSamplingInitial,
		SamplingThereafter: defaultSamplingThereafter,
	}
}

func createTracesExporter(ctx context.Context, set exporter.CreateSettings, config component.Config) (exporter.Traces, error) {
	cfg := config.(*Config)
	exporterLogger := createLogger(cfg, set.TelemetrySettings.Logger)
	debugExporter := newDebugExporter(exporterLogger, cfg.Verbosity)
	return exporterhelper.NewTracesExporter(ctx, set, config,
		debugExporter.pushTraces,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithShutdown(otlptext.LoggerSync(exporterLogger)),
	)
}

func createMetricsExporter(ctx context.Context, set exporter.CreateSettings, config component.Config) (exporter.Metrics, error) {
	cfg := config.(*Config)
	exporterLogger := createLogger(cfg, set.TelemetrySettings.Logger)
	debugExporter := newDebugExporter(exporterLogger, cfg.Verbosity)
	return exporterhelper.NewMetricsExporter(ctx, set, config,
		debugExporter.pushMetrics,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithShutdown(otlptext.LoggerSync(exporterLogger)),
	)
}

func createLogsExporter(ctx context.Context, set exporter.CreateSettings, config component.Config) (exporter.Logs, error) {
	cfg := config.(*Config)
	exporterLogger := createLogger(cfg, set.TelemetrySettings.Logger)
	debugExporter := newDebugExporter(exporterLogger, cfg.Verbosity)
	return exporterhelper.NewLogsExporter(ctx, set, config,
		debugExporter.pushLogs,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithShutdown(otlptext.LoggerSync(exporterLogger)),
	)
}

func createLogger(cfg *Config, logger *zap.Logger) *zap.Logger {
	core := zapcore.NewSamplerWithOptions(
		logger.Core(),
		1*time.Second,
		cfg.SamplingInitial,
		cfg.SamplingThereafter,
	)

	return zap.New(core)
}
