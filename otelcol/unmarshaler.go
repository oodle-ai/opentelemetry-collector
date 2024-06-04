// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcol // import "github.com/oodle-ai/opentelemetry-collector/otelcol"

import (
	"github.com/oodle-ai/opentelemetry-collector/confmap"
	"github.com/oodle-ai/opentelemetry-collector/connector"
	"github.com/oodle-ai/opentelemetry-collector/exporter"
	"github.com/oodle-ai/opentelemetry-collector/extension"
	"github.com/oodle-ai/opentelemetry-collector/otelcol/internal/configunmarshaler"
	"github.com/oodle-ai/opentelemetry-collector/processor"
	"github.com/oodle-ai/opentelemetry-collector/receiver"
	"github.com/oodle-ai/opentelemetry-collector/service"
	"github.com/oodle-ai/opentelemetry-collector/service/telemetry"
)

type configSettings struct {
	Receivers  *configunmarshaler.Configs[receiver.Factory]  `mapstructure:"receivers"`
	Processors *configunmarshaler.Configs[processor.Factory] `mapstructure:"processors"`
	Exporters  *configunmarshaler.Configs[exporter.Factory]  `mapstructure:"exporters"`
	Connectors *configunmarshaler.Configs[connector.Factory] `mapstructure:"connectors"`
	Extensions *configunmarshaler.Configs[extension.Factory] `mapstructure:"extensions"`
	Service    service.Config                                `mapstructure:"service"`
}

// unmarshal the configSettings from a confmap.Conf.
// After the config is unmarshalled, `Validate()` must be called to validate.
func unmarshal(v *confmap.Conf, factories Factories) (*configSettings, error) {

	telFactory := telemetry.NewFactory()
	defaultTelConfig := *telFactory.CreateDefaultConfig().(*telemetry.Config)

	// Unmarshal top level sections and validate.
	cfg := &configSettings{
		Receivers:  configunmarshaler.NewConfigs(factories.Receivers),
		Processors: configunmarshaler.NewConfigs(factories.Processors),
		Exporters:  configunmarshaler.NewConfigs(factories.Exporters),
		Connectors: configunmarshaler.NewConfigs(factories.Connectors),
		Extensions: configunmarshaler.NewConfigs(factories.Extensions),
		// TODO: Add a component.ServiceFactory to allow this to be defined by the Service.
		Service: service.Config{
			Telemetry: defaultTelConfig,
		},
	}

	return cfg, v.Unmarshal(&cfg)
}
