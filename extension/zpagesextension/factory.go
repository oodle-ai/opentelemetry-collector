// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zpagesextension // import "github.com/oodle-ai/opentelemetry-collector/extension/zpagesextension"

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/config/confignet"
	"github.com/oodle-ai/opentelemetry-collector/extension"
	"github.com/oodle-ai/opentelemetry-collector/extension/zpagesextension/internal/metadata"
)

const (
	defaultEndpoint = "localhost:55679"
)

// NewFactory creates a factory for Z-Pages extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(metadata.Type, createDefaultConfig, createExtension, metadata.ExtensionStability)
}

func createDefaultConfig() component.Config {
	return &Config{
		TCPAddr: confignet.TCPAddrConfig{
			Endpoint: defaultEndpoint,
		},
	}
}

// createExtension creates the extension based on this config.
func createExtension(_ context.Context, set extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	return newServer(cfg.(*Config), set.TelemetrySettings), nil
}
