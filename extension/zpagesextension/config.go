// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zpagesextension // import "github.com/oodle-ai/opentelemetry-collector/extension/zpagesextension"

import (
	"errors"

	"github.com/oodle-ai/opentelemetry-collector/component"
	"github.com/oodle-ai/opentelemetry-collector/config/confignet"
)

// Config has the configuration for the extension enabling the zPages extension.
type Config struct {
	// TCPAddr is the address and port in which the zPages will be listening to.
	// Use localhost:<port> to make it available only locally, or ":<port>" to
	// make it available on all network interfaces.
	TCPAddr confignet.TCPAddrConfig `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the extension configuration is valid
func (cfg *Config) Validate() error {
	if cfg.TCPAddr.Endpoint == "" {
		return errors.New("\"endpoint\" is required when using the \"zpages\" extension")
	}
	return nil
}
