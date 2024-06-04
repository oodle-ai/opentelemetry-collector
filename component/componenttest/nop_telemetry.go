// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package componenttest // import "github.com/oodle-ai/opentelemetry-collector/component/componenttest"

import (
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/config/configtelemetry"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

// NewNopTelemetrySettings returns a new nop telemetry settings for Create* functions.
func NewNopTelemetrySettings() component.TelemetrySettings {
	return component.TelemetrySettings{
		Logger:         zap.NewNop(),
		TracerProvider: nooptrace.NewTracerProvider(),
		MeterProvider:  noopmetric.NewMeterProvider(),
		MetricsLevel:   configtelemetry.LevelNone,
		Resource:       pcommon.NewResource(),
		ReportStatus: func(*component.StatusEvent) {
		},
	}
}
