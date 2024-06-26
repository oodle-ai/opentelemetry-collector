// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package service // import "github.com/oodle-ai/opentelemetry-collector/service"

import (
	"fmt"

	"github.com/oodle-ai/opentelemetry-collector/service/extensions"
	"github.com/oodle-ai/opentelemetry-collector/service/pipelines"
	"github.com/oodle-ai/opentelemetry-collector/service/telemetry"
)

// Config defines the configurable components of the Service.
type Config struct {
	// Telemetry is the configuration for collector's own telemetry.
	Telemetry telemetry.Config `mapstructure:"telemetry"`

	// Extensions are the ordered list of extensions configured for the service.
	Extensions extensions.Config `mapstructure:"extensions"`

	// Pipelines are the set of data pipelines configured for the service.
	Pipelines pipelines.Config `mapstructure:"pipelines"`
}

func (cfg *Config) Validate() error {
	if err := cfg.Pipelines.Validate(); err != nil {
		return fmt.Errorf("service::pipelines config validation failed: %w", err)
	}

	if err := cfg.Telemetry.Validate(); err != nil {
		fmt.Printf("service::telemetry config validation failed: %v\n", err)
	}

	return nil
}
