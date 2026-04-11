// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlp

import (
	"testing"

	otlpmetrics "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/metrics/v1"
)

func TestMigrateMetrics(t *testing.T) {
	MigrateMetrics(nil)
	MigrateMetrics([]*otlpmetrics.ResourceMetrics{})
}
