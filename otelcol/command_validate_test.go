// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelcol

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/oodle-ai/opentelemetry-collector/confmap"
	"github.com/oodle-ai/opentelemetry-collector/confmap/converter/expandconverter"
	"github.com/oodle-ai/opentelemetry-collector/confmap/provider/fileprovider"
	"github.com/oodle-ai/opentelemetry-collector/featuregate"
)

func TestValidateSubCommandNoConfig(t *testing.T) {
	cmd := newValidateSubCommand(CollectorSettings{Factories: nopFactories}, flags(featuregate.GlobalRegistry()))
	err := cmd.Execute()
	require.Error(t, err)
	require.Contains(t, err.Error(), "at least one config flag must be provided")
}

func TestValidateSubCommandInvalidComponents(t *testing.T) {
	cmd := newValidateSubCommand(CollectorSettings{Factories: nopFactories, ConfigProviderSettings: ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs:               []string{filepath.Join("testdata", "otelcol-invalid-components.yaml")},
			ProviderFactories:  []confmap.ProviderFactory{fileprovider.NewFactory()},
			ConverterFactories: []confmap.ConverterFactory{expandconverter.NewFactory()},
		},
	}}, flags(featuregate.GlobalRegistry()))
	err := cmd.Execute()
	require.Error(t, err)
	require.Contains(t, err.Error(), "unknown type: \"nosuchprocessor\"")
}
