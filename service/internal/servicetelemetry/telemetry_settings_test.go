// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package servicetelemetry

import (
	"testing"

	"github.com/stretchr/testify/require"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/config/configtelemetry"
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
	"github.com/oodle-ai/opentelemetry-collector/service/internal/status"
)

func TestSettings(t *testing.T) {
	set := TelemetrySettings{
		Logger:         zap.NewNop(),
		TracerProvider: nooptrace.NewTracerProvider(),
		MeterProvider:  noopmetric.NewMeterProvider(),
		MetricsLevel:   configtelemetry.LevelNone,
		Resource:       pcommon.NewResource(),
		Status: status.NewReporter(
			func(*component.InstanceID, *component.StatusEvent) {},
			func(err error) { require.NoError(t, err) }),
	}
	set.Status.Ready()
	set.Status.ReportStatus(
		&component.InstanceID{},
		component.NewStatusEvent(component.StatusStarting),
	)
	set.Status.ReportOKIfStarting(&component.InstanceID{})

	compSet := set.ToComponentTelemetrySettings(&component.InstanceID{})
	compSet.ReportStatus(component.NewStatusEvent(component.StatusStarting))
}
