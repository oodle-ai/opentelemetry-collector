// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package memorylimiterextension // import "github.com/oodle-ai/opentelemetry-collector/extension/memorylimiterextension"

//go:generate mdatagen metadata.yaml

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/extension"
	"github.com/oodle-ai/opentelemetry-collector/extension/memorylimiterextension/internal/metadata"
)

// NewFactory returns a new factory for the Memory Limiter extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability)
}

// CreateDefaultConfig creates the default configuration for extension. Notice
// that the default configuration is expected to fail for this extension.
func createDefaultConfig() component.Config {
	return &Config{}
}

func createExtension(_ context.Context, set extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	return newMemoryLimiter(cfg.(*Config), set.TelemetrySettings.Logger)
}
