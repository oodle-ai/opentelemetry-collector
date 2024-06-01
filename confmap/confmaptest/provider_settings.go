// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package confmaptest // import "github.com/oodle-ai/opentelemetry-collector/confmap/confmaptest"

import (
	"go.uber.org/zap"

	"github.com/oodle-ai/opentelemetry-collector/confmap"
)

func NewNopProviderSettings() confmap.ProviderSettings {
	return confmap.ProviderSettings{Logger: zap.NewNop()}
}
