// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package httpprovider // import "github.com/oodle-ai/opentelemetry-collector/confmap/provider/httpprovider"

import (
	"github.com/oodle-ai/opentelemetry-collector/confmap"
	"github.com/oodle-ai/opentelemetry-collector/confmap/provider/internal/configurablehttpprovider"
)

// NewFactory returns a factory for a confmap.Provider that reads the configuration from a http server.
//
// This Provider supports "http" scheme.
//
// One example for HTTP URI is: http://localhost:3333/getConfig
func NewFactory() confmap.ProviderFactory {
	return confmap.NewProviderFactory(newProvider)
}

func newProvider(set confmap.ProviderSettings) confmap.Provider {
	return configurablehttpprovider.New(configurablehttpprovider.HTTPScheme, set)
}
