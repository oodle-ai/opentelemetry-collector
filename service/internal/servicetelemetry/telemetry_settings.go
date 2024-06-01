// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package servicetelemetry // import "github.com/oodle-ai/opentelemetry-collector/service/internal/servicetelemetry"

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/config/configtelemetry"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
	"github.com/oodle-ai/opentelemetry-collector/service/internal/status"
)

// TelemetrySettings mirrors component.TelemetrySettings except for the mechanism for reporting
// status. Service-level status reporting has additional methods which can report status for
// components by their InstanceID whereas the component versions are tied to a specific component.
type TelemetrySettings struct {
	// Logger that the factory can use during creation and can pass to the created
	// component to be used later as well.
	Logger *zap.Logger

	// TracerProvider that the factory can pass to other instrumented third-party libraries.
	TracerProvider trace.TracerProvider

	// MeterProvider that the factory can pass to other instrumented third-party libraries.
	MeterProvider metric.MeterProvider

	// MetricsLevel controls the level of detail for metrics emitted by the collector.
	// Experimental: *NOTE* this field is experimental and may be changed or removed.
	MetricsLevel configtelemetry.Level

	// Resource contains the resource attributes for the collector's telemetry.
	Resource pcommon.Resource

	// Status contains a Reporter that allows the service to report status on behalf of a
	// component.
	Status *status.Reporter
}

// ToComponentTelemetrySettings returns a TelemetrySettings for a specific component derived from
// this service level Settings object.
func (s TelemetrySettings) ToComponentTelemetrySettings(id *component.InstanceID) component.TelemetrySettings {
	statusFunc := status.NewReportStatusFunc(id, s.Status.ReportStatus)
	return component.TelemetrySettings{
		Logger:         s.Logger,
		TracerProvider: s.TracerProvider,
		MeterProvider:  s.MeterProvider,
		MetricsLevel:   s.MetricsLevel,
		Resource:       s.Resource,
		ReportStatus:   statusFunc,
	}
}
