// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcoltest // import "github.com/oodle-ai/opentelemetry-collector/otelcol/otelcoltest"

import (
	"context"

	"github.com/oodle-ai/opentelemetry-collector/confmap"
	"github.com/oodle-ai/opentelemetry-collector/confmap/converter/expandconverter"
	"github.com/oodle-ai/opentelemetry-collector/confmap/provider/envprovider"
	"github.com/oodle-ai/opentelemetry-collector/confmap/provider/fileprovider"
	"github.com/oodle-ai/opentelemetry-collector/confmap/provider/httpprovider"
	"github.com/oodle-ai/opentelemetry-collector/confmap/provider/yamlprovider"
	"github.com/oodle-ai/opentelemetry-collector/otelcol"
)

// LoadConfig loads a config.Config  from file, and does NOT validate the configuration.
func LoadConfig(fileName string, factories otelcol.Factories) (*otelcol.Config, error) {
	// Read yaml config from file
	provider, err := otelcol.NewConfigProvider(otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs: []string{fileName},
			ProviderFactories: []confmap.ProviderFactory{
				fileprovider.NewFactory(),
				envprovider.NewFactory(),
				yamlprovider.NewFactory(),
				httpprovider.NewFactory(),
			},
			ConverterFactories: []confmap.ConverterFactory{expandconverter.NewFactory()},
		},
	})
	if err != nil {
		return nil, err
	}
	return provider.Get(context.Background(), factories)
}

// LoadConfigAndValidate loads a config from the file, and validates the configuration.
func LoadConfigAndValidate(fileName string, factories otelcol.Factories) (*otelcol.Config, error) {
	cfg, err := LoadConfig(fileName, factories)
	if err != nil {
		return nil, err
	}
	return cfg, cfg.Validate()
}
