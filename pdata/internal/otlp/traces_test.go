// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlp

import (
	"testing"

	otlptrace "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/trace/v1"
)

func TestMigrateTraces(t *testing.T) {
	MigrateTraces(nil)
	MigrateTraces([]*otlptrace.ResourceSpans{})
}
