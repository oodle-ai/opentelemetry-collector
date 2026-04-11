// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlp

import (
	"testing"

	otlplogs "github.com/oodle-ai/opentelemetry-collector/pdata/internal/data/protogen/logs/v1"
)

func TestMigrateLogs(t *testing.T) {
	MigrateLogs(nil)
	MigrateLogs([]*otlplogs.ResourceLogs{})
}
