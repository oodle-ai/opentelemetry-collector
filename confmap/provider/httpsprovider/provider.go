// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package httpsprovider // import "github.com/oodle-ai/opentelemetry-collector/confmap/provider/httpsprovider"

import (
	"github.com/oodle-ai/opentelemetry-collector/confmap"
	"github.com/oodle-ai/opentelemetry-collector/confmap/provider/internal/configurablehttpprovider"
)

// NewFactory returns a factory for a confmap.Provider that reads the configuration from a https server.
//
// This Provider supports "https" scheme. One example of an HTTPS URI is: https://localhost:3333/getConfig
//
// To add extra CA certificates you need to install certificates in the system pool. This procedure is operating system
// dependent. E.g.: on Linux please refer to the `update-ca-trust` command.
func NewFactory() confmap.ProviderFactory {
	return confmap.NewProviderFactory(newProvider)
}

func newProvider(set confmap.ProviderSettings) confmap.Provider {
	return configurablehttpprovider.New(configurablehttpprovider.HTTPSScheme, set)
}
